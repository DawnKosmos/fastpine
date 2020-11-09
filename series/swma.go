package series

import "github.com/dawnkosmos/fastpine/cist"

type swma struct {
	USR
	src Series

	data       *cist.Cist
	tempResult float64
}

/*
	pine_swma(x) =>
	x[3] * 1 / 6 + x[2] * 2 / 6 + x[1] * 2 / 6 + x[0] * 1 / 6
	plot(pine_swma(close))

*/

//Swma is quivalent to swma(src)
func Swma(src Series) Series {
	var s swma
	s.src = src
	s.resolution = src.Resolution()
	s.starttime = src.Starttime() + int64(s.resolution*4)
	if src.UpdateGroup() != nil {
		s.ug = src.UpdateGroup()
		s.ug.Add(&s)
	}

	s.data = cist.New()
	f := src.Data()
	var fOut []float64 = make([]float64, 0, len(f))
	x3, x2, x1, x0 := 1/6*f[0], 2/6*f[1], 2/6*f[2], 1/6*f[3]
	fOut = append(fOut, x0+x1+x2+x3)

	for i := 3; i < len(f); i++ {
		x3 = x2 / 2
		x2 = x1
		x1 = x0 * 2
		x0 = f[i] * 1 / 6
		fOut = append(fOut, x0+x1+x2+x3)
	}

	s.data.InitData(fOut)

	return &s
}

func (s *swma) Add() {
	s.data.Add()
}

func (s *swma) Update() {
	_ = "kk"
}

func (s *swma) Value(index int) float64 {
	return s.data.Get(index)
}

func (s *swma) Data() []float64 {
	return s.data.GetData()
}
