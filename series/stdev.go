package series

import (
	"math"

	"github.com/dawnkosmos/fastpine/cist"
)

type stdev struct {
	src Series
	l   int

	USR

	data *cist.Cist
}

/*
variance(samples):
  M := 0
  S := 0
  for k from 1 to N:
    x := samples[k]
    oldM := M
    M := M + (x-M)/k
    S := S + (x-M)*(x-oldM)
  return S/(N-1)


Mk-1 = Mk - (xk - Mk) / (k - 1).

Sk-1 = Sk - (xk – Mk-1) * (xk – Mk).

*/

//Stdev is equivalent to stdev(src, len)
func Stdev(src Series, l int) Series {
	var s stdev
	s.src = src
	s.resolution = src.Resolution()
	s.starttime = src.Starttime() + int64(s.resolution*l)

	if src.UpdateGroup() != nil {
		s.ug = src.UpdateGroup()
		(*s.ug).Add(&s)
	}
	s.data = cist.New()
	f := src.Data()
	fOut := make([]float64, 0, len(f))

	//init

	var m, oldM, S float64
	var m1, s1 float64

	for i := 0; i < l; i++ {
		x := f[i]
		oldM = m
		m = m + (x-m)/float64(i+1)
		S += (x - m) * (x - oldM)
	}

	fOut = append(fOut, math.Sqrt(S/float64(l)))

	lf := float64(l)
	for i := l; i < len(f); i++ {
		x1 := f[i-l]
		m1 = (lf*m - x1) / (lf - 1)
		s1 = S - (x1-m1)*(x1-m)

		x := f[i]
		m = m1 + (x-m1)/lf
		S = s1 + (x-m)*(x-m1)

		//fmt.Println(S, s1)
		fOut = append(fOut, math.Sqrt(S/lf))
	}

	s.data.InitData(fOut)
	return &s

}

func (s *stdev) Update() {
	_ = "kek"
}

func (s *stdev) Add() {
	s.data.Add()
}

func (s *stdev) Value(index int) float64 {
	return (*s.data).Get(index)
}

func (s *stdev) Data() []float64 {
	return (*s.data).GetData()
}
