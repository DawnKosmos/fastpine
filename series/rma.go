package series

import (
	"github.com/dawnkosmos/fastpine/cist"
	"github.com/dawnkosmos/fastpine/helper"
)

type rma struct {
	src Series
	len int
	USR

	data       *cist.Cist
	alpha      float64
	tempResult float64
}

//Rma is equivalent to rma(src,len)
func Rma(src Series, l int) Series {
	var s rma

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
	s.alpha = 1 / float64(l)

	avg := helper.FloatAverage(f[:l])
	r := make([]float64, 0, len(f))
	r = append(r, avg)
	for i := l; i < len(f); i++ {
		avg = f[i]*s.alpha + (1-s.alpha)*avg
		r = append(r, avg)
	}

	s.data.InitData(r)

	return &s
}

func (s *rma) Add() {
	s.data.Add()
}

func (s *rma) Update() {
	_ = "kek"
}

func (r *rma) Value(index int) float64 {
	return (*r.data).Get(index)
}

func (r *rma) Data() []float64 {
	return (*r.data).GetData()
}
