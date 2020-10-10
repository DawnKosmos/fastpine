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
	//SubscribeTicker([]string, chan Ticker)
	OHCLV(ticker string, resolution int, start int64, end int64) ([]Candle, error)
	Name() string
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
func ConvertCandleResolution(c ...Candle) Candle {
	var out Candle = Candle{0, 0, c[0].Low, c[0].Open, 0, c[0].StartTime}
	for _, i := range c {
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

func ConvertChartResolution(new, old int64, ch []Candle) ([]Candle, error) {
	if new == old {
		return ch, nil
	}
	var startingPoint, endPoint int
	if old > new || new%old != 0 {
		return ch, fmt.Errorf("New Res %v and old %v do not fit", new, old)
	}
	for i, c := range ch {
		if c.StartTime.Unix()%new == 0 {
			startingPoint = i
			break
		}
	}
	for i := len(ch) - 1; i != 1; i-- {
		if ch[i].StartTime.Unix()%new == 0 {
			endPoint = i
			break
		}
	}
	multiplicator := int(new / old)

	var out []Candle

	for i := startingPoint; i < endPoint; {
		out = append(out, ConvertCandleResolution(ch[i:i+multiplicator]...))
		i = i + multiplicator
	}
	out = append(out, ConvertCandleResolution(ch[endPoint:len(ch)]...))

	return out, nil
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
