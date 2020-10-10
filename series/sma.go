package series

import (
	"github.com/dawnkosmos/fastpine/cist"
	"github.com/dawnkosmos/fastpine/exchange"
)

type SMA struct {
	src        Series
	len        int
	starttime  int64
	resolution int
	ug         *exchange.UpdateGroup
	/*Cist Element saves:
	Source/alpha
	Result
	*/
	data       *cist.Cist
	tempResult float64
}

/* calculation
SMA = avg(src, len)

*/

//Sma creates an
func Sma(src Series, l int) *SMA {
	var s SMA
	//Init Values
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
	l1 := len(*f)
	//r is the results array which gets written in the s.data List
	var r []float64 = make([]float64, 0, l1)
	//Calc the avg of the first Len candles
	avg := Average((*f)[:l])
	r = append(r, avg)
	alpha := 1 / float64(l)
	//Iterrating while dynamical calculating the values and add them
	for i := l; i < len(*f); i++ {
		avg = avg - (*f)[i-l]*alpha + (*f)[i]*alpha
		r = append(r, avg)
	}
	l2 := len(r)
	s.data.InitData(&r)
	c := make([]float64, 0, l)
	for _, v := range (*f)[l1-l:] {
		c = append(c, v/alpha)
	}
	//Fills the temp saved elements which help to calculate the sma faster
	s.data.FillElements(l, c, r[l2-l:])
	return &s
}

func (s *SMA) Update() {
	v := s.src.Value(0)
	if v == s.tempResult {
		return
	}
	s.tempResult = v
	//Can be faster by using temp data instead of always calculating
	/*Cist Element saves:
	Source/alpha
	Result
	*/
	smaF, smaL := s.data.First(), s.data.Last()
	result := smaF[1] - smaL[1]
	b := v / float64(s.len)
	s.data.Update(result+b, b, result+b)
}

func (s *SMA) Value(index int) float64 {
	return (*s.data).Get(index)
}

func (s *SMA) Resolution() int {
	return s.resolution
}

func (s *SMA) Starttime() int64 {
	return s.starttime
}

func (s *SMA) Data() *[]float64 {
	return (*s.data).GetData()
}

func (s *SMA) UpdateGroup() *exchange.UpdateGroup {
	return s.ug
}

func (s *SMA) Add() {
	s.data.Add()
}
