package ftx

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/buger/jsonparser"
	"github.com/gorilla/websocket"
)

//Most of the subscription code is inspired by the "github.com/go-numb/go-ftx/" package. I changed some stuff and added more comments

//a request for a WS subscription
type requestWS struct {
	Op      string `json:"op"`
	Channel string `json:"channel"`
	Market  string `json:"market"`
}

//Ticker gives you actual marketing data which than can be used to calculated the market price.
//Which is the median of {bid,ask,last}
type Ticker struct {
	Bid     float64   `json:"bid"`
	Ask     float64   `json:"ask"`
	BidSize float64   `json:"bidSize"`
	AskSize float64   `json:"askSize"`
	Last    float64   `json:"last"`
	Time    time.Time `json:"time"`
}

//Fill is a private subscription. A channel where we get a signal when we got a fill
type Fill struct {
	Fee         float64 `json:"fee"`
	FeeCurrency string  `json:"feeCurrency"`
	Future      string  `json:"future"`
	FeeRate     float64 `json:"feeRate"`
	Market      string  `json:"market"`
	Type        string  `json:"type"`
	Liquidity   string  `json:"liquidity"`

	BaseCurrency  string `json:"baseCurrency"`
	QuoteCurrency string `json:"quoteCurrency"`

	Side  string    `json:"side"`
	Price float64   `json:"price"`
	Size  float64   `json:"size"`
	Time  time.Time `json:"time"`

	ID      int `json:"id"`
	OrderID int `json:"orderId"`
	TradeID int `json:"tradeId"`
}

const (
	UNDEFINED = iota
	ERROR
	TICKER
	TRADES
	ORDERBOOK
	ORDERS
	FILLS
)

type Orderbook struct {
	Bids [][]float64 `json:"bids"`
	Asks [][]float64 `json:"asks"`
	// Action return update/partial
	Action   string    `json:"action"`
	Time     time.Time `json:"time"`
	Checksum int       `json:"checksum"`
}

type ResponseWS struct {
	Type   int
	Symbol string

	Ticker    Ticker
	Trades    []Trade
	Orderbook Orderbook
	Orders    Order
	Fills     Fill
	Results   error
}

// {"op": "login", "args": {"key": "<api_key>", "sign": "<signature>", "time": 1111}}
type requestForPrivateWS struct {
	Op   string                 `json:"op"`
	Args map[string]interface{} `json:"args"`
}

/*
Connect gives you a subscription for the given channels and symbols you inserted
Parameters:
	conn: A websocket connection to the FTX ws
	ctx: context
	ch: In this channel the asked subsprictions get saved
	channels: Channels e.g. Orderbook, trades, ticker, fills
	symbols: ticker e.g. "BTC-PERP", "ETH-PERP"
*/
func (f *FTX) Connect(ctx context.Context, ch chan ResponseWS, channels, symbols []string, l *log.Logger) error {
	if l == nil {
		l = log.New(os.Stdout, "ftx websocket", log.Llongfile)
	}
	var err error
	//This creates a websocked connection to the FTX ws Server
	/*conn, _, err := websocket.DefaultDialer.Dial("wss://ftx.com/ws/", nil)
	if err != nil {
		return err
	}*/

	f.conn, _, err = websocket.DefaultDialer.Dial("wss://ftx.com/ws/", nil)

	if err != nil {
		return err
	}

	if f.key != "" {
		//sign up to private connection
		if err := signatureWS(f.conn, f.key, f.secret, f.UsedAccount); err != nil {
			return err
		}
	}

	//Here we send a message to the FTX ws Server telling which Subscriptionschannels from which Ticker we like to describe
	// An example would be channels = {"Ticker"} symbols = {"BTC-PERP","ETH-PERP"};
	// FTX now starts to send the Ticker of btc and eth to us
	if err := Subscribe(f.conn, channels, symbols); err != nil {
		return err
	}

	//Pings every 15 seconds FTX, so they know that we are still online
	//If you don't ping FTX, they assume that we disconnected and stop sending us messages
	//So be aware that if you lose internet for some time you have to restart the connections.
	go ping(f.conn)

	//In this goroutine we reseive FTX subscriptions Messages, Parse them and add them to the ResponseWS channel
	go func() {
		defer f.conn.Close()
		defer Unsubscribe(f.conn, channels, symbols)

	RESTART:
		for {
			var res ResponseWS
			//the msg we get from FTX. Has this information in it.
			/*
				channel: A subscription e.g. "Ticker"
				market:	market name e.g. "BTC-PERP"
				type: The type of message. e.g. "Subscribed", "unsubscribed", "error", "Update"
				code (optional)
				msg (optional)
				data(optional): the data the channel provides e.g. ticker provides information about  {bid,ask,last,time}

			*/
			_, msg, err := f.conn.ReadMessage()
			if err != nil {
				l.Printf("[ERROR]: msg error: %+v", err)
				res.Type = ERROR
				res.Results = fmt.Errorf("%v", err)
				ch <- res
				break RESTART
			}

			typeMsg, err := jsonparser.GetString(msg, "type")
			if typeMsg == "error" {
				l.Printf("[ERROR]: error: %+v", string(msg))
				res.Type = ERROR
				res.Results = fmt.Errorf("%v", string(msg))
				ch <- res
				break RESTART
			}

			//Channel which we subscribed
			channel, err := jsonparser.GetString(msg, "channel")
			if err != nil {
				l.Printf("[ERROR]: channel error: %+v", string(msg))
				res.Type = ERROR
				res.Results = fmt.Errorf("%v", string(msg))
				ch <- res
				break RESTART
			}

			//Pair we subscriped
			market, err := jsonparser.GetString(msg, "market")
			if err != nil {
				l.Printf("[ERROR]: market err: %+v", string(msg))
				res.Type = ERROR
				res.Results = fmt.Errorf("%v", string(msg))
				ch <- res
				break RESTART
			}

			res.Symbol = market

			//Data that got sent gets received and saved as a string
			data, _, _, err := jsonparser.Get(msg, "data")
			if err != nil {
				if typeMsg == "subscribed" {
					l.Printf("[SUCCESS]: %s %+v", typeMsg, string(msg))
					continue
				} else {
					err = fmt.Errorf("[ERROR]: data err: %v %s", err, string(msg))
					l.Println(err)
					res.Type = ERROR
					res.Results = err
					ch <- res
					break RESTART
				}
			}

			//Depending on which kind of subscrition we reseive we parse the data which saved in a string differently into a ResponseWS
			switch channel {
			case "ticker":
				res.Type = TICKER
				if err := json.Unmarshal(data, &res.Ticker); err != nil {
					l.Printf("[WARN]: cant unmarshal ticker %+v", err)
					continue
				}

			case "trades":
				res.Type = TRADES
				if err := json.Unmarshal(data, &res.Trades); err != nil {
					l.Printf("[WARN]: cant unmarshal trades %+v", err)
					continue
				}

			case "orderbook":
				res.Type = ORDERBOOK
				if err := json.Unmarshal(data, &res.Orderbook); err != nil {
					l.Printf("[WARN]: cant unmarshal orderbook %+v", err)
					continue
				}

			default:
				res.Type = UNDEFINED
				res.Results = fmt.Errorf("%v", string(msg))
			}

			//We send the parsed data in a channel which we can use in other programms.
			ch <- res

		}
	}()
	return nil
}

func signatureWS(conn *websocket.Conn, key string, secret []byte, subaccount string) error {
	// key: your API key
	// time: integer current timestamp (in milliseconds)
	// sign: SHA256 HMAC of the following string, using your API secret: <time>websocket_login
	// subaccount: (optional) subaccount name
	// As an example, if:

	// time: 1557246346499
	// secret: 'Y2QTHI23f23f23jfjas23f23To0RfUwX3H42fvN-'
	// sign would be d10b5a67a1a941ae9463a60b285ae845cdeac1b11edc7da9977bef0228b96de9

	// One websocket connection may be logged in to at most one user. If the connection is already authenticated, further attempts to log in will result in 400s.

	msec := time.Now().UTC().UnixNano() / int64(time.Millisecond)

	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(fmt.Sprintf("%dwebsocket_login", msec)))
	args := map[string]interface{}{
		"key":  key,
		"sign": hex.EncodeToString(mac.Sum(nil)),
		"time": msec,
	}
	if "" != subaccount {
		args["subaccount"] = subaccount
	}

	if err := conn.WriteJSON(&requestForPrivateWS{
		Op:   "login",
		Args: args,
	}); err != nil {
		return err
	}

	return nil
}

//subscribe subscibes to a channel
//you can use this function to add further subscriptions after your conn is already connected to the ws server
func Subscribe(conn *websocket.Conn, channels, symbols []string) error {
	if conn == nil {
		return fmt.Errorf("Cant subsribe if no connection")
	}

	if symbols != nil {
		for i := range channels {
			for j := range symbols {
				if err := conn.WriteJSON(&requestWS{
					Op:      "subscribe",
					Channel: channels[i],
					Market:  symbols[j],
				}); err != nil {
					return err
				}
			}
		}
	} else {
		for i := range channels {
			if err := conn.WriteJSON(&requestWS{
				Op:      "subscribe",
				Channel: channels[i],
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

//unsubscribe to the Given Channels(Subscriptions) & Symbols
func Unsubscribe(conn *websocket.Conn, channels, symbols []string) error {
	if conn == nil {
		return fmt.Errorf("Cant unsubscribe if no connection")
	}

	if symbols != nil {
		for i := range channels {
			for j := range symbols {
				if err := conn.WriteJSON(&requestWS{
					Op:      "unsubscribe",
					Channel: channels[i],
					Market:  symbols[j],
				}); err != nil {
					return err
				}
			}
		}
	} else {
		for i := range channels {
			if err := conn.WriteJSON(&requestWS{
				Op:      "unsubscribe",
				Channel: channels[i],
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

//ping is pinging FTX every 15 seconds, to not disconnect
func ping(conn *websocket.Conn) (err error) {
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := conn.WriteMessage(websocket.PingMessage, []byte(`{"op": "pong"}`)); err != nil {
				goto EXIT
			}
		}
	}
EXIT:
	return err
}
