package series

import (
	"github.com/dawnkosmos/fastpine/cist"
	"github.com/dawnkosmos/fastpine/exchange"
)

type EMA struct {
	src        Series
	len        int
	starttime  int64
	resolution int
	ug         *exchange.UpdateGroup

	/*Cist Element saves:
	Just last value
	*/
	alpha      float64
	data       *cist.Cist
	tempResult float64
}

/* EMA Calculation
1. Calculate the SMA

(Period Values / Number of Periods)
2. Calculate the Multiplier

(2 / (Number of Periods + 1) therefore (2 / (5+1) = 33.333%
3. Calculate the EMA
For the first EMA, we use the SMA(previous day) instead of EMA(previous day).

EMA = {Close - EMA(previous day)} x multiplier + EMA(previous day)
*/

func Ema(src Series, l int) *EMA {
	var s EMA
	s.len = l
	s.src = src
	s.resolution = src.Resolution()
	s.starttime = src.Starttime() + int64(s.resolution*l)
	if src.UpdateGroup() != nil {
		s.ug = src.UpdateGroup()
		(*s.ug).Add(&s)
	}
	s.alpha = 2.0 / float64(l+1)
	s.data = cist.New()
	f := src.Data()
	l1 := len(*f)
	r := make([]float64, 0, l1)
	avg := Average((*f)[:l])
	r = append(r, avg)
	for i := l; i < l1; i++ {
		avg = ((*f)[i]-avg)*s.alpha + avg
		r = append(r, avg)
	}
	s.data.InitData(&r)
	return &s

}

func (s *EMA) Update() {
	v := s.src.Value(0)
	if v == s.tempResult {
		return
	}
	s.tempResult = v
	r := (v-s.data.Get(1))*s.alpha + s.data.Get(1)
	s.data.Update(r)
}

func (s *EMA) UpdateGroup() *exchange.UpdateGroup {
	return s.ug
}

func (s *EMA) Value(index int) float64 {
	return (*s.data).Get(index)
}

func (s *EMA) Data() *[]float64 {
	return (*s.data).GetData()
}

func (s *EMA) Add() {
	s.data.Add()
}

func (s *EMA) Starttime() int64 {
	return s.starttime
}

func (s *EMA) Resolution() int {
	return s.resolution
}
