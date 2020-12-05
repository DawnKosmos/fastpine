package series

import "github.com/dawnkosmos/fastpine/exchange"

//source as SERIES

type Inputtype uint

const (
	OPEN Inputtype = iota
	HIGH
	LOW
	CLOSE
	VOLUME
)

type CandlestickProvider interface {
	StarttimeResolutionUpdategroup
	Value(i int) exchange.Candle
	Data() []exchange.Candle
}

//source Represents all the Candlestick information
type source struct {
	o   CandlestickProvider
	Get func(e exchange.Candle) float64
	d   []float64
	USR
}

//Source gets you specific ohcl data
// Source(CLOSE, o) gets you close value, also stuff like OHCVL4 or HCL3, OC2 got implemented
func Source(In Inputtype, o CandlestickProvider) Series {
	var s source
	s.o = o
	s.starttime = o.Starttime()
	s.resolution = o.Resolution()
	s.ug = o.UpdateGroup()
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
	}

	s.getData()
	return &s
}

func (s *source) Value(i int) float64 {
	c := s.o.Value(i)
	return s.Get(c)
}

func (s *source) getData() {
	var out []float64
	for _, c := range s.o.Data() {
		out = append(out, s.Get(c))
	}
	s.d = out
}

func (s *source) Data() []float64 {
	return s.d
}
