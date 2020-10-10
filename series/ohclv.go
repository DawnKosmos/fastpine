package series

import (
	"time"

	"github.com/dawnkosmos/fastpine/exchange"
)

const (
	OPEN Inputtype = iota
	HIGH
	LOW
	CLOSE
	VOLUME
	HL2
	OC2
)

type Inputtype uint

type OHCLV struct {
	E          *exchange.Exchange
	Ticker     string
	Starttime  int64
	Resolution int
	lenChart   int
	ch         []exchange.Candle
	Ug         *exchange.UpdateGroup
}

func NewOHCLV(e exchange.Exchange, ticker string, startime int64, resolution int, Ug *exchange.UpdateGroup) *OHCLV {
	var o OHCLV
	o.E = &e
	o.Ticker = ticker
	o.Resolution = resolution
	o.Starttime = startime
	o.ch, _ = (*o.E).OHCLV(ticker, o.Resolution, o.Starttime, time.Now().Unix())
	o.lenChart = len(o.ch)
	return &o

}

func (o *OHCLV) Update() {
	c, _ := (*o.E).Actual(o.Ticker, int64(o.Resolution))
	o.ch[o.lenChart] = c
}

func (o *OHCLV) Add() {
	c, _ := (*o.E).Actual(o.Ticker, int64(o.Resolution))
	o.ch = append(o.ch, c)
	o.lenChart++
}

func (o *OHCLV) Value(i int) exchange.Candle {
	return o.ch[o.lenChart-i-1]
}

//SOURCE as SERIES

type SOURCE struct {
	o   *OHCLV
	Get func(e exchange.Candle) float64
	d   *[]float64
	ug  *exchange.UpdateGroup
}

func GetSRC(In Inputtype, o *OHCLV) *SOURCE {
	var s SOURCE
	s.o = o
	switch In {
	case OPEN:
		s.Get = func(e exchange.Candle) float64 {
			return e.Open
		}
	case HIGH:
		s.Get = func(e exchange.Candle) float64 {
			return e.High
		}
	case LOW:
		s.Get = func(e exchange.Candle) float64 {
			return e.Low
		}
	case CLOSE:
		s.Get = func(e exchange.Candle) float64 {
			return e.Close
		}
	case VOLUME:
		s.Get = func(e exchange.Candle) float64 {
			return e.Volume
		}
	case HL2:
		s.Get = func(e exchange.Candle) float64 {
			return (e.High + e.Low) / 2
		}
	case OC2:
		s.Get = func(e exchange.Candle) float64 {
			return (e.Open + e.Close) / 2
		}
	}

	s.getData()
	return &s
}

func (s *SOURCE) Value(i int) float64 {
	c := (*s.o).Value(i)
	return s.Get(c)
}

func (s *SOURCE) Resolution() int {
	return (*s.o).Resolution
}

func (s *SOURCE) Starttime() int64 {
	return (*s.o).Starttime
}

func (s *SOURCE) getData() {
	var out []float64
	for _, c := range (*s.o).ch {
		out = append(out, s.Get(c))
	}
	s.d = &out
}

func (s *SOURCE) Data() *[]float64 {
	return s.d
}

func (s *SOURCE) UpdateGroup() *exchange.UpdateGroup {
	return s.ug
}
