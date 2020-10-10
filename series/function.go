package series

import (
	"github.com/dawnkosmos/fastpine/exchange"
)

type FUNCTION struct {
	src       []Series
	op        func(index int, v ...Series) float64
	starttime int64
	ug        *exchange.UpdateGroup
	fOut      *[]float64
	delay     int
	len       int
	fBool     *[]bool
}

func Function(operator func(i int, v ...Series) float64, delay int, src ...Series) *FUNCTION {
	var s FUNCTION
	s.src = src
	s.op = operator
	s.ug = src[0].UpdateGroup()
	s.starttime, s.len = getStarttime(src...)
	s.starttime = int64(delay*s.src[0].Resolution()) + s.starttime
	s.delay = delay
	var f []Series = make([]Series, 0, len(s.src))
	var iterLen int = s.len - s.delay
	for _, v := range s.src {
		f = append(f, fakeSeries(v, &iterLen))
	}
	//Delay offsets starttime
	fOut := make([]float64, 0, s.len)
	temp := iterLen
	for i := 0; i < temp; i++ {
		iterLen--
		fOut = append(fOut, s.op(0, f...))
	}
	s.fOut = &fOut
	return &s
}

//Return starttime, len of the shortest array
func getStarttime(src ...Series) (int64, int) {
	s := src[0].Starttime()
	n := 0
	for i := 1; i < len(src); i++ {
		b := bigger(s, src[i].Starttime())
		if s != b {
			s = b
			n = i
		}
	}
	return s, len(*src[n].Data())
}

func (s *FUNCTION) Starttime() int64 {
	return s.starttime
}

func (s *FUNCTION) Resolution() int {
	return s.src[0].Resolution()
}

func (s *FUNCTION) Value(index int) float64 {
	return s.op(index, s.src...)
}

func (s *FUNCTION) ValueB(index int) bool {
	return s.op(index, s.src...) > 0
}

func (s *FUNCTION) DataB() *[]bool {
	if s.fBool == nil {
		f := make([]bool, 0, len(*s.fOut))
		for _, v := range *s.fOut {
			f = append(f, v > 0.0)
		}
		s.fBool = &f
	}
	return s.fBool
}

func (s *FUNCTION) UpdateGroup() *exchange.UpdateGroup {
	return s.ug
}

func (s *FUNCTION) Data() *[]float64 {
	return s.fOut
}

//fSeries is a (Fake)series needed to dynamical init a Series based on a function
type fSeries struct {
	src Series
	len *int
}

//TODO Using delay as int pointer makes it little bit faster so we dont always have to increment all delays

func fakeSeries(src Series, len *int) *fSeries {
	var s fSeries = fSeries{src, len}
	return &s
}

func (s *fSeries) Value(index int) float64 {
	return s.src.Value(index + *s.len)
	//return float64(index + s.delay)
}

func (s *fSeries) Data() *[]float64 {
	return nil
}

func (s *fSeries) Resolution() int {
	return 0
}

func (s *fSeries) Starttime() int64 {
	return 0
}

func (s *fSeries) UpdateGroup() *exchange.UpdateGroup {
	return nil
}
