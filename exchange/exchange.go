package exchange

import (
	"fmt"
	"time"
)

type Value interface{}

type Instance struct {
	E         *Exchange
	Frequenzy time.Ticker
	U         []*Updater
}

const HOUR int = 3600

type Exchange interface {
	Name() string
	//SubscribeTicker([]string, chan Ticker)

	//OHCLV kriegt noch ein Bool für LiveData
	OHCLV(ticker string, resolution int, start int64, end int64) ([]Candle, error)
	Actual(Ticker string, resolution int64) (Candle, error)
	//SubscribeFills([]string, chan Fill)
}

type Ticker struct {
	Exchange string    `json:"exchange"`
	Ticker   string    `json:"ticker"`
	Ask      float64   `json:"ask"`
	Bid      float64   `json:"bid"`
	Last     float64   `json:"last"`
	Time     time.Time `json:"time"`
}

type Fill struct {
	Exchange string    `json:"exchange"`
	Ticker   string    `json:"ticker"`
	Price    float64   `json:"price"`
	Amount   float64   `json:"amount"`
	Time     time.Time `json:"time"`
}

type Candle struct {
	Close     float64   `json:"close"`
	High      float64   `json:"high"`
	Low       float64   `json:"low"`
	Open      float64   `json:"open"`
	Volume    float64   `json:"volume"`
	StartTime time.Time `json:"startTime"`
}

func (c Candle) OHCL4() float64 {
	return (c.Open + c.Close + c.High + c.Low) / 4
}

//HELP FUNCTIONS
//Getting starting timestamps
func DateToTime(day, month, year string) int64 {
	t, _ := time.Parse("01 02 2006 15:04 MST", month+" "+day+" "+year+" 00:00 GMT")
	return t.Unix()
}

func DateToTimeHourly(day, month, year, hour string) int64 {
	t, _ := time.Parse("01 02 2006 15:04 MST", month+" "+day+" "+year+" "+hour+" GMT")
	return t.Unix()
}

//ConvertResolution converts the a lower resolution into a higher resolution
func ConvertCandleResolution(c []Candle) Candle {
	var out Candle = Candle{c[0].Close, c[0].High, c[0].Low, c[0].Open, c[0].Volume, c[0].StartTime}
	for _, i := range c[1:] {
		out.Close = i.Close
		out.Volume += i.Volume
		if i.High > out.High {
			out.High = i.High
		}
		if i.Low < out.Low {
			out.Low = i.Low
		}
	}
	return out
}

//New res must me greater than old
func ConvertChartResolution(oldResolution, newResolution int64, ch []Candle) ([]Candle, error) {
	if newResolution == oldResolution {
		return ch, nil
	}

	if oldResolution > newResolution || newResolution%oldResolution != 0 {
		return ch, fmt.Errorf("New Res %v and old %v do not fit", newResolution, oldResolution)
	}

	quotient := int(newResolution / oldResolution)

	var newChart []Candle = make([]Candle, 0, len(ch)/quotient)

	for _, c := range ch {
		if c.StartTime.Unix()%newResolution != 0 {
			ch = ch[1:]
		} else {
			break
		}
	}

	for {
		if len(ch) < quotient {
			break
		}
		newChart = append(newChart, ConvertCandleResolution(ch[:quotient]))
		ch = ch[quotient:]
	}
	if len(ch) != 0 {
		newChart = append(newChart, ConvertCandleResolution(ch))
	}

	return newChart, nil
}

func median(a, b, c float64) float64 {
	if a > b {
		if a < c {
			return a
		} else if b > c {
			return b
		} else {
			return c
		}
	} else {
		if a > c {
			return a
		} else if b < c {
			return b
		} else {
			return c
		}
	}
}

/*
func getChart(exchange string, ticker string, resolution int) (chart []Candle, doesExist bool) {
	fileName := exchange + ticker + strconv.Itoa(resolution) + ".txt"

}

func checkSrcFolder() error {
	_, err := os.Stat("src")
	if err != nil {
		err := os.Mkdir("src", 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
*/
