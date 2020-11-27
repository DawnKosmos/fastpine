package series

import "github.com/dawnkosmos/fastpine/cist"

type transForm struct {
	src           Series
	oldResolution int

	USR

	data *cist.Cist
}

/* TODO
func MatchToTimeFrame(src Series, newResolution int) Series{
	var s transForm
	s.src = src
	s.resolution = newResolution
	s.oldResolution = src.Resolution()
	sOff := Offset(src,1)
	s.starttime = sOff.Starttime()
	if s.oldResolution%newResolution != 0{
		return src
	}

	if src.UpdateGroup() != nil {
		s.ug = src.UpdateGroup()
		(*s.ug).Add(&s)
	}

	muliplier = s.oldResolution / newResolution
}



func (s *fisher) Update() {
	_ = "kek"
}

func (s *fisher) Add() {
	s.data.Add()
}

func (s *fisher) Value(index int) float64 {
	return (*s.data).Get(index)
}

func (s *fisher) Data() []float64 {
	return (*s.data).GetData()
}
*/
