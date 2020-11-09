package series

import (
	"github.com/dawnkosmos/fastpine/cist"
	"github.com/dawnkosmos/fastpine/helper"
)

type sma struct {
	src Series
	len int
	USR
	/*Cist Element saves:
	Source/alpha
	Result
	*/
	data       *cist.Cist
	tempResult float64
}

/* calculation
sma = avg(src, len)

*/

//Sma is equivalent to sma(src,len)
func Sma(src Series, l int) Series {
	var s sma
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
	l1 := len(f)
	//r is the results array which gets written in the s.data List
	var r []float64 = make([]float64, 0, l1)
	//Calc the avg of the first Len candles
	avg := helper.FloatAverage(f[:l])
	r = append(r, avg)
	alpha := 1 / float64(l)
	//Iterrating while dynamical calculating the values and add them
	for i := l; i < len(f); i++ {
		avg = avg - f[i-l]*alpha + f[i]*alpha
		r = append(r, avg)
	}
	l2 := len(r)
	s.data.InitData(r)
	c := make([]float64, 0, l)
	for _, v := range f[l1-l:] {
		c = append(c, v/alpha)
	}
	//Fills the temp saved elements which help to calculate the sma faster
	s.data.FillElements(c, r[l2-l:])
	return &s
}

func (s *sma) Update() {
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

func (s *sma) Value(index int) float64 {
	return (*s.data).Get(index)
}

func (s *sma) Data() []float64 {
	return (*s.data).GetData()
}

func (s *sma) Add() {
	s.data.Add()
}
