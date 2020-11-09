package series

import (
	"github.com/dawnkosmos/fastpine/cist"
	"github.com/dawnkosmos/fastpine/helper"
)

type sum struct {
	src Series
	len int
	USR

	data       *cist.Cist
	tempResult float64
}

//Sum is quivalent to sum(src, len)
func Sum(src Series, l int) Series {
	var s sum
	s.src = src
	s.resolution = src.Resolution()
	s.starttime = src.Starttime() + int64(s.resolution*l)

	if src.UpdateGroup() != nil {
		s.ug = src.UpdateGroup()
		(*s.ug).Add(&s)
	}

	s.data = cist.New()
	f := src.Data()
	l1 := len(f)
	var fOut []float64 = make([]float64, 0, l1)
	//init
	initSum := helper.FloatSum(f[:l])
	fOut = append(fOut, initSum)

	for i := l; i < len(f); i++ {
		initSum = initSum - f[i-l] + f[i]
		fOut = append(fOut, initSum)
	}

	s.data.InitData(fOut)

	return &s
}

func (s *sum) Value(index int) float64 {
	return (*s.data).Get(index)
}

func (s *sum) Data() []float64 {
	return (*s.data).GetData()
}

func (s *sum) Add() {
	s.data.Add()
}

func (s *sum) Update() {
	_ = "kek"
}
