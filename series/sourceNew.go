package series

import (
	"github.com/dawnkosmos/fastpine/exchange"
)

type nource struct {
	o  CandlestickProvider
	op func(e exchange.Candle) float64
	d  []float64
	USR
}

func Nource(o CandlestickProvider, op func(e exchange.Candle) float64) Series {
	var s nource
	s.o = o
	s.op = op
	s.starttime = o.Starttime()
	s.resolution = o.Resolution()
	s.ug = o.UpdateGroup()

	var out []float64 = make([]float64, 0, len(o.Data()))

	for _, c := range o.Data() {
		out = append(out, s.op(c))
	}
	s.d = out

	return &s
}

func (s *nource) Value(i int) float64 {
	c := s.o.Value(i)
	return s.op(c)
}

func (s *nource) Data() []float64 {
	return s.d
}

func Open(c CandlestickProvider) Series {
	fn := func(e exchange.Candle) float64 {
		return e.Open
	}
	return Nource(c, fn)
}

func Close(c CandlestickProvider) Series {
	fn := func(e exchange.Candle) float64 {
		return e.Close
	}
	return Nource(c, fn)
}

func High(c CandlestickProvider) Series {
	fn := func(e exchange.Candle) float64 {
		return e.High
	}
	return Nource(c, fn)
}

func Low(c CandlestickProvider) Series {
	fn := func(e exchange.Candle) float64 {
		return e.Low
	}
	return Nource(c, fn)
}

func Volume(c CandlestickProvider) Series {
	fn := func(e exchange.Candle) float64 {
		return e.Volume
	}
	return Nource(c, fn)
}

func HL2(c CandlestickProvider) Series {
	fn := func(e exchange.Candle) float64 {
		return (e.High + e.Low) / 2
	}
	return Nource(c, fn)
}

func OHCL4(c CandlestickProvider) Series {
	fn := func(e exchange.Candle) float64 {
		return (e.Open + e.Close + e.High + e.Low) / 4
	}
	return Nource(c, fn)
}

func OC2(c CandlestickProvider) Series {
	fn := func(e exchange.Candle) float64 {
		return (e.Open + e.Close) / 2
	}
	return Nource(c, fn)
}

func HCL3(c CandlestickProvider) Series {
	fn := func(e exchange.Candle) float64 {
		return (e.Close + e.High + e.Low) / 3
	}
	return Nource(c, fn)
}
