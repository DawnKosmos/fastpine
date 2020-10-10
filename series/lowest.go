package series

import (
	"github.com/dawnkosmos/fastpine/cist"
	"github.com/dawnkosmos/fastpine/exchange"
)

type LOWEST struct {
	src Series
	len int
	ug  *exchange.UpdateGroup

	data *cist.Cist
	//TEMP VALUES
	tempResult      float64
	lowPosition     int
	pastlowPosition int
	pastLow         float64
}

func Lowest(src Series, l int) *LOWEST {
	var s LOWEST
	s.len = l
	s.src = src
	f := src.Data()
	l1 := len(*f)
	s.data = cist.New()
	var r []float64 = make([]float64, 0, l1)
	var low float64 = (*f)[0]
	var lowPosition int = 0
	r = append(r, low)
	for i := 1; i < l; i++ {
		if (*f)[i] <= low {
			lowPosition = i
			low = (*f)[i]
		}
		r = append(r, low)
	}

	for i := l; i < l1; i++ {
		if lowPosition < i-l {
			lowPosition = getLowest((*f)[lowPosition+1:i]) + lowPosition
			low = (*f)[lowPosition]
		}
		if (*f)[i] <= low {
			lowPosition = i
			low = (*f)[i]
		}
		r = append(r, low)
	}

	s.data.InitData(&r)
	s.data.FillElements(l)

	return &s
}

func (s *LOWEST) Update() {
	v := s.src.Value(0)
	if v == s.tempResult {
		return
	}
	s.tempResult = v
	if s.pastLow > v {
		s.data.Update(v)
	}

}

func getLowest(f []float64) int {
	low := f[0]
	lowPosition := 0
	for i := 1; i < len(f); i++ {
		if f[i] <= low {
			lowPosition = i
			low = f[i]
		}
	}
	return lowPosition
}

func (s *LOWEST) Resolution() int {
	return s.src.Resolution()
}

func (s *LOWEST) Starttime() int64 {
	return s.src.Starttime()
}

func (s *LOWEST) UpdateGroup() *exchange.UpdateGroup {
	return s.ug
}
