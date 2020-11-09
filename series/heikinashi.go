package series

import (
	"math"

	"github.com/dawnkosmos/fastpine/exchange"
)

type heikinashi struct {
	src *OHCLV

	USR
	ch []exchange.Candle
}

/*
HeikinAshi turns an inputed chart into a heikin ashi candlesticks chart.
Open = (open of previous bar + close of previous bar)/2
Close = (open + high + low + close)/4
High = the maximum value from the high, open, or close of the current period
Low = the minimum value from the low, open, or close of the current period
*/
func HeikinAshi(src *OHCLV) *heikinashi {
	var s heikinashi
	s.src = src
	s.starttime = src.Starttime()
	s.resolution = src.Resolution()
	if src.UpdateGroup() != nil {
		s.ug = src.UpdateGroup()
		(*s.ug).Add(&s)
	}
	s.ch = make([]exchange.Candle, 0, len(src.ch)+5)
	var c exchange.Candle
	c.Open = src.ch[0].Open
	c.Close = src.ch[0].OHCL4()
	c.Low = math.Min(src.ch[0].Open, c.Open)
	c.High = math.Max(src.ch[0].High, c.Open)
	c.Volume = src.ch[0].Volume
	c.StartTime = src.ch[0].StartTime
	s.ch = append(s.ch, c)
	for _, v := range src.ch[1:] {
		c.Open = (c.Open + c.Close) / 2
		c.Close = v.OHCL4()
		c.Low = math.Min(v.Open, c.Open)
		c.High = math.Min(v.Open, c.Open)
		c.Volume = v.Volume
		c.StartTime = v.StartTime
		s.ch = append(s.ch, c)
	}

	return &s
}

func (o *heikinashi) Value(i int) exchange.Candle {
	return o.ch[len(o.ch)-i-1]
}

func (o *heikinashi) Add() {
	c := o.src.Value(0)
	ha := o.Value(0)
	close := c.OHCL4()
	open := (ha.Open + ha.Close) / 2
	o.ch = append(o.ch, exchange.Candle{
		Open:      open,
		Close:     close,
		Low:       math.Min(c.Low, open),
		High:      math.Max(c.High, open),
		Volume:    c.Volume,
		StartTime: c.StartTime,
	})
}

func (o *heikinashi) Update() {
	c := o.src.Value(0)
	ha := &o.ch[len(o.ch)-1]
	(*ha).Close = ha.OHCL4()
	(*ha).Low = math.Min(c.Low, (*ha).Open)
	(*ha).High = math.Max(c.High, (*ha).Open)
	(*ha).Volume = c.Volume
}

func (o *heikinashi) Data() []exchange.Candle {
	return o.ch
}
