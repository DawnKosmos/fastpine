package series

import (
	"github.com/dawnkosmos/fastpine/cist"
	"github.com/dawnkosmos/fastpine/exchange"
)

type HIGHEST struct {
	src Series
	len int
	ug  *exchange.UpdateGroup

	data *cist.Cist
	//TEMP VALUES
	tempResult       float64
	highPosition     int
	pasthighPosition int
	pastHigh         float64
}

func Highest(src Series, l int) *HIGHEST {
	var s HIGHEST
	s.len = l
	s.src = src
	f := src.Data()
	l1 := len(*f)
	s.data = cist.New()
	var r []float64 = make([]float64, 0, l1)
	var high float64 = (*f)[0]
	var highPosition int = 0
	r = append(r, high)
	for i := 1; i < l; i++ {
		if (*f)[i] >= high {
			highPosition = i
			high = (*f)[i]
		}
		r = append(r, high)
	}

	for i := l; i < l1; i++ {
		if highPosition < i-l {
			highPosition = getHighest((*f)[highPosition+1:i]) + highPosition
			high = (*f)[highPosition]
		}
		if (*f)[i] >= high {
			highPosition = i
			high = (*f)[i]
		}
		r = append(r, high)
	}

	s.data.InitData(&r)
	s.data.FillElements(l)

	return &s
}

func (s *HIGHEST) Update() {
	v := s.src.Value(0)
	if v == s.tempResult {
		return
	}
	s.tempResult = v
	if s.pastHigh < v {
		s.data.Update(v)
	}

}

func getHighest(f []float64) int {
	high := f[0]
	highPosition := 0
	for i := 1; i < len(f); i++ {
		if f[i] >= high {
			highPosition = i
			high = f[i]
		}
	}
	return highPosition
}

func (s *HIGHEST) Resolution() int {
	return s.src.Resolution()
}

func (s *HIGHEST) Starttime() int64 {
	return s.src.Starttime()
}

func (s *HIGHEST) UpdateGroup() *exchange.UpdateGroup {
	return s.ug
}
