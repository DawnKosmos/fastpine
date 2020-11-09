package series

import (
	"github.com/dawnkosmos/fastpine/exchange"
)

type function struct {
	src       []Series
	op        func(index int, v ...Series) float64
	starttime int64
	ug        *exchange.UpdateGroup
	fOut      []float64
	delay     int
	len       int
	fBool     []bool
}

/*Function allows you to implement complicated indicators
Sadly its really complicated and slow, Should be used really carefully
Usually you should be able to implement everything with the other given indicators,
You also can write me and I will add your specific indicator.

Function implements the Series and Condition Interface

Look at the end of  this this file to see some example implementation
*/
func Function(operator func(index int, v ...Series) float64, delay int, src ...Series) *function {
	var s function
	//Init
	s.src = src
	s.op = operator
	s.ug = src[0].UpdateGroup()
	s.starttime, s.len = getStarttime(src...)
	s.starttime = int64(delay*s.src[0].Resolution()) + s.starttime
	s.delay = delay
	//Creating a Series Slice later filled with fSeries
	var f []Series = make([]Series, 0, len(s.src))
	var iterLen int = s.len - s.delay
	for _, v := range s.src {
		f = append(f, fakeSeries(v, &iterLen))
	}

	fOut := make([]float64, 0, s.len)
	temp := iterLen
	//We calculate the Values in this For loop
	for i := 0; i < temp; i++ {
		iterLen--
		fOut = append(fOut, s.op(0, f...))
	}
	s.fOut = fOut
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
	return s, len(src[n].Data())
}

func (s *function) Resolution() int {
	return s.src[0].Resolution()
}

func (s *function) Value(index int) float64 {
	return s.op(index, s.src...)
}

func (s *function) Starttime() int64 {
	return s.starttime
}

func (s *function) ValueB(index int) bool {
	return s.op(index, s.src...) > 0
}

func (s *function) DataB() []bool {
	if s.fBool == nil {
		f := make([]bool, 0, len(s.fOut))
		for _, v := range s.fOut {
			f = append(f, v > 0.0)
		}
		s.fBool = f
	}
	return s.fBool
}

func (s *function) UpdateGroup() *exchange.UpdateGroup {
	return s.ug
}

func (s *function) Data() []float64 {
	return s.fOut
}

//fSeries is a (Fake)series needed to dynamical init a Series based on a function
type fSeries struct {
	src Series
	len *int
}

func fakeSeries(src Series, len *int) *fSeries {
	var s fSeries = fSeries{src, len}
	return &s
}

func (s *fSeries) Value(index int) float64 {
	return s.src.Value(index + *s.len)
}

//Empty functions but needed for the Interface
func (s *fSeries) Data() []float64 { return nil }

func (s *fSeries) Resolution() int { return 0 }

func (s *fSeries) Starttime() int64 { return 0 }

func (s *fSeries) UpdateGroup() *exchange.UpdateGroup { return nil }

//EXAMPLE
//Lol TODO :P
