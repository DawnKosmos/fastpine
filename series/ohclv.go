package series

import (
	"time"

	"github.com/dawnkosmos/fastpine/exchange"
)

type OHCLV struct {
	E      *exchange.Exchange
	Ticker string

	lenChart int
	ch       []exchange.Candle
	USR
}

//NewOHCLV gets you your candlestick chart
func NewOHCLV(e exchange.Exchange, ticker string, startime int64, resolution int, Ug *exchange.UpdateGroup) *OHCLV {
	var o OHCLV
	o.E = &e
	o.Ticker = ticker
	o.resolution = resolution
	o.starttime = startime
	o.ch, _ = (*o.E).OHCLV(ticker, o.resolution, o.starttime, time.Now().Unix())
	o.lenChart = len(o.ch)
	return &o

}

func (o *OHCLV) Update() {
	c, _ := (*o.E).Actual(o.Ticker, int64(o.resolution))
	o.ch[o.lenChart-1] = c
}

func (o *OHCLV) Add() {
	c, _ := (*o.E).Actual(o.Ticker, int64(o.resolution))
	o.ch = append(o.ch, c)
	o.lenChart++
}

func (o *OHCLV) Value(i int) exchange.Candle {
	return o.ch[o.lenChart-i-1]
}

func (o *OHCLV) Data() []exchange.Candle {
	return o.ch
}
