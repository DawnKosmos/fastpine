package series

import (
	"github.com/dawnkosmos/fastpine/cist"
	"github.com/dawnkosmos/fastpine/exchange"
)

/*Calculation
100 * (Current Close - Lowest Low) / (Highest High - Lowest Low)
*/
type STOCH struct {
	src  Series
	high Series
	low  Series
	len  int

	starttime  int64
	resolution int

	data *cist.Cist
	ug   *exchange.UpdateGroup
}

func Stoch(src Series, highS Series, lowS Series, le int) *STOCH {
	var s STOCH
	//Init
	s.src, s.high, s.low, s.len = src, highS, lowS, le
	s.resolution = src.Resolution()
	b := bigger(lowS.Starttime(), highS.Starttime())
	s.starttime = bigger(b, src.Starttime())
	if src.UpdateGroup() != nil {
		s.ug = src.UpdateGroup()
		(*s.ug).Add(&s)
	}
	s.data = cist.New()
	f := (*src.Data())[(src.Starttime()-s.starttime)/int64(s.resolution):]
	h := (*highS.Data())[(highS.Starttime()-s.starttime)/int64(s.resolution):]
	l := (*lowS.Data())[(lowS.Starttime()-s.starttime)/int64(s.resolution):]
	var r []float64 = make([]float64, 0, len(f)+10)
	var highPosition, lowPosition int = 0, 0
	var high, low float64 = h[0], l[0]
	r = append(r, 100*(f[0]-low)/(high-low))
	//calculate the first le amount of elements
	for i := 1; i < le; i++ {
		if h[i] >= high {
			highPosition = i
			high = h[i]
		}
		if l[i] <= low {
			lowPosition = i
			low = l[i]
		}
		r = append(r, 100*(f[i]-low)/(high-low))
	}
	for i := le; i < len(f); i++ {
		if lowPosition < i-le {
			lowPosition = getLowest(l[lowPosition+1:i]) + lowPosition
			low = l[lowPosition]
		}

		if l[i] <= low {
			lowPosition = i
			low = l[i]
		}

		if highPosition < i-le {
			highPosition = getHighest(h[highPosition+1:i]) + highPosition
			high = h[highPosition]
		}

		if h[i] >= high {
			highPosition = i
			high = h[i]
		}

		r = append(r, 100*(f[i]-low)/(high-low))
	}
	l1 := len(f)
	s.data.InitData(&r)
	s.data.FillElements(le, f[l1-le:], h[l1-le:], l[l1-le:])
	return &s

}

//100 * (Current Close - Lowest Low) / (Highest High - Lowest Low)

func (s *STOCH) Update() {
	_ = "kek"
}

func (s *STOCH) Add() {
	s.data.Add()
}

func (s *STOCH) UpdateGroup() *exchange.UpdateGroup {
	return s.ug
}

func (s *STOCH) Value(index int) float64 {
	return (*s.data).Get(index)
}

func (s *STOCH) Resolution() int {
	return s.resolution
}

func (s *STOCH) Starttime() int64 {
	return s.starttime
}

func (s *STOCH) Data() *[]float64 {
	return (*s.data).GetData()
}
