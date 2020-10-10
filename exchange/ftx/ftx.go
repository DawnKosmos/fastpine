package ftx

import (
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

/*
	Most of the FTX code is inspired by the FTX and go-numb/ftx library. Props to them.
	I did not include them fully for performance reasons and also some adjustments were needed or else the performance would have suffered significant
	I also added comments to some stuff, so not experienced programmers like me, understand the code better :)
*/

type FTX struct {
	Client *http.Client
	//KEY & SECRET ARE FROM MAIN
	key    string
	secret []byte
	//If you want to have access another subaccount you have to change UseAccount. select no name for main, else use name of your subaccount
	UsedAccount string
	Subaccounts []string
	conn        *websocket.Conn
	//TickerSubscription []string
	//ch                 chan realtime.Response
}

const URL = "https://ftx.com/api/"

func New(api string, secret string, UsedAccount string, subaccounts []string) *FTX {
	return &FTX{&http.Client{}, api, []byte(secret), url.PathEscape(UsedAccount), pathEscape(subaccounts), nil}
}

/*

 */

//pathEscape corrects if a Subaccountname has special characters or Empty Space in its name
func pathEscape(s []string) []string {
	for _, a := range s {
		a = url.PathEscape(a)
	}
	return s
}

func (client *FTX) Name() string {
	return "FTX"
}
