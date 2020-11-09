package series

import "github.com/dawnkosmos/fastpine/cist"

//ROC = [(CurrentClose - Close n periods ago) / (Close n periods ago)] X 100
type roc struct {
	src Series
	len int
	USR

	data       *cist.Cist
	tempResult float64
}

//Roc is the equivalent to roc(src, len)
func Roc(src Series, l int) Series {
	var s roc

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

	var r []float64 = make([]float64, 0, l1)

	for i := l; i < len(f); i++ {
		r = append(r, (f[i]-f[i-l])/f[i-l])
	}

	s.data.InitData(r)

	return &s
}

func (s *roc) Value(index int) float64 {
	return (*s.data).Get(index)
}

func (s *roc) Data() []float64 {
	return (*s.data).GetData()
}

func (s *roc) Add() {
	s.data.Add()
}

func (s *roc) Update() {
	_ = "kek"
}
