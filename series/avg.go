package series

import (
	"github.com/dawnkosmos/fastpine/cist"
	"github.com/dawnkosmos/fastpine/helper"
)

type avg struct {
	src []Series

	USR

	data *cist.Cist
}

//Avg returns the average value of the inputed Series
func Avg(src ...Series) Series {
	var s avg
	s.src = src
	var f [][]float64
	for _, v := range src {
		f = append(f, v.Data())
	}

	if src[0].UpdateGroup() != nil {
		s.ug = src[0].UpdateGroup()
		(*s.ug).Add(&s)
	}

	shortest := helper.FloatArrLowestLen(f...)
	for i, v := range f {
		f[i] = v[len(v)-shortest:]
	}

	s.data = cist.New()
	fOut := make([]float64, 0, shortest)

	var sum float64
	for i := 0; i < shortest; i++ {
		sum = 0
		for j := 0; j < len(f); i++ {
			sum = sum + f[i][j]
		}
		fOut = append(fOut, sum/float64(len(f)))
	}
	s.data.InitData(fOut)

	return &s
}

func (s *avg) Value(index int) float64 {
	return (*s.data).Get(index)
}

func (s *avg) Data() []float64 {
	return (*s.data).GetData()
}

func (s *avg) Update() {
	_ = "kek"
}

func (s *avg) Add() {
	s.data.Add()
}
