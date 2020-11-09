package series

import (
	"github.com/dawnkosmos/fastpine/cist"
	"github.com/dawnkosmos/fastpine/exchange"
)

type lowest struct {
	src Series
	len int
	USR

	data *cist.Cist
	//TEMP VALUES
	tempResult float64
}

//Lowest is equivalent to lowest(src, len)
func Lowest(src Series, l int) Series {
	var s lowest
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
	lo := s.data.Lowest(0)
	lowe := lo.Get(0)
	fOut = append(fOut, lowe)

	var e *cist.Element
	for _, v := range f[l:] {
		e = s.data.NILO(v)

		if v <= lowe {
			lowe = v
			lo = s.data.GetEle(0)
		}
		if e == lo {
			lo = s.data.Lowest(0)
			lowe = lo.Get(0)
		}
		fOut = append(fOut, lowe)
	}

	s.data.InitData(fOut)

	return &s
}

func (s *lowest) Resolution() int {
	return s.src.Resolution()
}

func (s *lowest) Starttime() int64 {
	return s.src.Starttime()
}

func (s *lowest) UpdateGroup() *exchange.UpdateGroup {
	return s.ug
}

func (s *lowest) Data() []float64 {
	return s.data.GetData()
}

func (s *lowest) Value(index int) float64 {
	return s.data.Get(index)
}

func (s *lowest) Update() {
	_ = "kek"
}

func (s *lowest) Add() {
	s.data.Add()
}
