package series

import (
	"fmt"
	"math"

	"github.com/dawnkosmos/fastpine/cist"
)

type fisher struct {
	src Series
	len int

	USR

	data *cist.Cist
}

//TODO

func Fisher(src Series, l int) Series {
	var s fisher
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
	//var val1, val2, fish float64
	//Init  Values
	s.data.FillElements(f[:l])
	lo, hi := s.data.Lowest(0), s.data.Highest(0)
	lowest, highest := lo.Get(0), hi.Get(0)
	var fish float64 = 0.0
	var value1, value2 float64

	value1 = fishCalc1(f[l-1], lowest, highest, value1)

	value2 = fishCalc2(value1)
	fmt.Println(value1, value2, lowest, highest)
	fish = fishCalc(value2, fish)

	fOut = append(fOut, fish)
	var e *cist.Element
	for _, v := range f[l:] {
		e = s.data.NILO(v)

		if v >= highest {
			highest = v
			hi = s.data.GetEle(0)
		}
		if e == hi {
			hi = s.data.Highest(0)
			highest = hi.Get(0)
		}

		if v <= lowest {
			lowest = v
			lo = s.data.GetEle(0)
		}
		if e == lo {
			lo = s.data.Lowest(0)
			lowest = lo.Get(0)
		}

		value1 = fishCalc1(v, lowest, highest, value1)
		fish = fishCalc(fishCalc2(value1), fish)

		fOut = append(fOut, fish)
	}

	s.data.InitData(fOut)
	return &s

}

func fishCalc1(hl2, lowest, highest, prev float64) float64 {
	return 0.66*((hl2-lowest)/(highest-lowest)-0.5) + 0.67*prev
}

func fishCalc2(value1 float64) (v float64) {
	if value1 > 0.999 {
		v = 0.9999
	} else {
		if value1 < -0.999 {
			v = -0.9999
		} else {
			v = value1
		}
	}
	return
}

func fishCalc(value2 float64, fish float64) float64 {
	return 0.5*math.Log((1+value2)/(1-value2)) + 0.5*fish
}

/*
highest = Highest(hl2, Length)
lowest = Lowest(hl2,Length)
nValue1 = 0.33 * 2 * ((hl2 - lowest) / (highest - lowest) - 0.5) + 0.67 * nz(nValue1[1])
nValue2 = iff(nValue1 > .99,  .999,
	        iff(nValue1 < -.99, -.999, nValue1))
nFish = 0.5 * log((1 + nValue2) / (1 - nValue2)) + 0.5 * nz(nFish[1])
*/

func (s *fisher) Update() {
	_ = "kek"
}

func (s *fisher) Add() {
	s.data.Add()
}

func (s *fisher) Value(index int) float64 {
	return (*s.data).Get(index)
}

func (s *fisher) Data() []float64 {
	return (*s.data).GetData()
}
