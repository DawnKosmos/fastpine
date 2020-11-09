package series

import (
	"github.com/dawnkosmos/fastpine/cist"
	"github.com/dawnkosmos/fastpine/exchange"
)

type highest struct {
	src Series
	len int

	USR

	data *cist.Cist
	//TEMP VALUES
	tempResult float64
}

//Highest is quivalent to highest(src, len)
func Highest(src Series, l int) Series {
	var s highest
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
	var fOut []float64 = make([]float64, 0, len(f))
	s.data.FillElements(f[:l])

	hi := s.data.Highest(0)
	high := hi.Get(0)

	fOut = append(fOut, high)

	var e *cist.Element
	for _, v := range f[l:] {
		e = s.data.NILO(v)

		if v >= high {
			high = v
			hi = s.data.GetEle(0)
		}
		if e == hi {
			hi = s.data.Highest(0)
			high = hi.Get(0)
		}

		fOut = append(fOut, high)
	}

	s.data.InitData(fOut)
	return &s
}

func (s *highest) Update() {
	_ = "kek"
}

func (s *highest) Add() {
	_ = "kek"
}

func (s *highest) Resolution() int {
	return s.src.Resolution()
}

func (s *highest) Starttime() int64 {
	return s.src.Starttime()
}

func (s *highest) UpdateGroup() *exchange.UpdateGroup {
	return s.ug
}

func (s *highest) Data() []float64 {
	return s.data.GetData()
}

func (s *highest) Value(index int) float64 {
	return s.data.Get(index)
}
