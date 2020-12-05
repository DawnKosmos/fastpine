package series

import (
	"github.com/dawnkosmos/fastpine/cist"
)

type td struct {
	src Series
	len int

	USR

	data *cist.Cist
}

/*
TD = close > close[4] ?nz(TD[1])+1:0
TS = close < close[4] ?nz(TS[1])+1:0

TDUp = TD - valuewhen(TD < TD[1], TD , 1 )
TDDn = TS - valuewhen(TS < TS[1], TS , 1 )
*/

func TD(src Series, l int) Series {
	var s td
	s.len = l
	s.src = src

	s.resolution = src.Resolution()
	s.starttime = src.Starttime() + int64(s.resolution*l)
	if src.UpdateGroup() != nil {
		s.ug = src.UpdateGroup()
		(*s.ug).Add(&s)
	}
	s.data = cist.New()
	f := src.Data()

	out := make([]float64, 0, len(f))
	var val float64

	for i, v := range f[l:] {
		if f[i] < v {
			if val < 0 {
				val = 1.0
			} else {
				val += 1.0
			}
		} else {
			if val > 0 {
				val = -1.0
			} else {
				val -= 1.0
			}
		}
		out = append(out, val)
	}

	s.data.InitData(out)

	return &s
}

func (s *td) Update() {
	_ = "kek"
}

func (s *td) Add() {
	s.data.Add()
}

func (s *td) Value(index int) float64 {
	return (*s.data).Get(index)
}

func (s *td) Data() []float64 {
	return (*s.data).GetData()
}
