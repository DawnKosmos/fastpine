package series

import (
	"github.com/dawnkosmos/fastpine/cist"
	"github.com/dawnkosmos/fastpine/helper"
)

type TR struct {
	close, high, low Series

	USR

	data *cist.Cist

	c1, h1, l1 float64
}

//trueRange = na(high[1])? high-low : max(max(high - low, abs(high - close[1])), abs(low - close[1]))
func Truerange(close, high, low Series) *TR {
	var s TR
	s.close, s.high, s.low = close, high, low

	s.resolution = close.Resolution()
	s.starttime = close.Starttime()

	c, h, l := close.Data(), high.Data(), low.Data()

	r := make([]float64, 0, len(c))
	r = append(r, h[0]-l[0])

	var hlasb, abslc float64

	for i := 1; i < len(c); i++ {
		hlasb = helper.FloatMax(h[i]-l[i], helper.FloatAbs(h[i]-c[i-1]))
		abslc = helper.FloatAbs(l[i] - c[i-1])
		r = append(r, helper.FloatMax(hlasb, abslc))
	}
	s.data = cist.New()
	s.data.InitData(r)

	s.c1, s.h1, s.l1 = c[len(c)-2], h[len(h)-2], l[len(l)-2]

	return &s

}

func (s *TR) Value(index int) float64 {
	return s.data.Get(index)
}

func (s *TR) Data() []float64 {
	return s.data.GetData()
}

func (s *TR) Add() {
	s.data.Add()
}

func (s *TR) Update() {
	_ = "kek"
}
