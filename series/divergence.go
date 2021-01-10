package series

import (
	"github.com/dawnkosmos/fastpine/cist"
	"github.com/dawnkosmos/fastpine/helper"
)

/*
ind > ind[1] and clo < clo[1]

hbullDiv(ind,clo)=>
    ind < ind[1 ] and clo > clo[1]

strongDiv(ind,clo)=>
ind > ind[1] and clo > clo[1]

weakDiv(ind,clo)=>
	ind < ind[1] and clo < clo[1]

	var ind = 0.0
	var cl = 0.0

	if c1
		ind:=ccc[1]
		cl := haLow[1]

	var indS = 0.0
	var clS = 0.0

	if c2
		indS := ccc[1]
		clS := haHigh[1]

	bullDiv(ind, clo)=>
*/

type divergence struct {
	src       Series
	indicator Series
	trigger   Condition

	USR

	data *cist.Cist
}

/*
Div
HiddenDiv
BearishCon
BullishCon

*/

func LongDiv(src Series, indicator Series, trigger Condition) Series {
	var s divergence
	s.src, s.indicator, s.trigger = src, indicator, trigger

	s.src.Resolution()
	s.starttime = bigger(src.Starttime(), indicator.Starttime(), trigger.Starttime())

	if src.UpdateGroup() != nil {
		s.ug = src.UpdateGroup()
		(*s.ug).Add(&s)
	}
	s.data = cist.New()
	cl, ind, c1 := Lowest(s.src, 2).Data(), Offset(s.indicator, 1).Data(), s.trigger.DataB()
	var tcl, tind float64

	shortest := helper.Min(len(cl), len(ind), len(c1))
	cl = cl[len(cl)-shortest:]
	ind = ind[len(ind)-shortest:]
	c1 = c1[len(c1)-shortest:]

	out := make([]float64, 0, len(cl))

	for i, v := range c1 {
		if v {
			out = append(out, divHelp(cl[i], ind[i], tcl, tind))
			tcl = cl[i]
			tind = ind[i]
		} else {
			out = append(out, 0.0)
		}
	}
	s.data.InitData(out)

	return &s

}

func (s *divergence) Update() {
	_ = "kek"
}

func (s *divergence) Add() {
	s.data.Add()
}

func divHelp(cl, ind, cl1, ind1 float64) float64 {
	if cl >= cl1 {
		if ind >= ind1 {
			return 4.0
		} else {
			return 2.0
		}
	} else {
		if ind >= ind1 {
			return 1.0
		} else {
			return 3.0
		}
	}
}

func (s *divergence) Value(index int) float64 {
	return (*s.data).Get(index)
}

func (s *divergence) Data() []float64 {
	return (*s.data).GetData()
}

func ShortDiv(src Series, indicator Series, trigger Condition) Series {
	var s divergence
	s.src, s.indicator, s.trigger = src, indicator, trigger

	s.src.Resolution()
	s.starttime = bigger(src.Starttime(), indicator.Starttime(), trigger.Starttime())

	if src.UpdateGroup() != nil {
		s.ug = src.UpdateGroup()
		(*s.ug).Add(&s)
	}
	s.data = cist.New()
	cl, ind, c1 := Highest(s.src, 2).Data(), Offset(s.indicator, 1).Data(), s.trigger.DataB()
	var tcl, tind float64

	shortest := helper.Min(len(cl), len(ind), len(c1))
	cl = cl[len(cl)-shortest:]
	ind = ind[len(ind)-shortest:]
	c1 = c1[len(c1)-shortest:]

	out := make([]float64, 0, len(cl))

	for i, v := range c1 {
		if v {
			out = append(out, divHelp(cl[i], ind[i], tcl, tind))
			tcl = cl[i]
			tind = ind[i]
		} else {
			out = append(out, 0.0)
		}
	}
	s.data.InitData(out)

	return &s
}

func divHelp2(cl, ind, cl1, ind1 float64) float64 {
	if cl >= cl1 {
		if ind >= ind1 {
			return 4.0
		} else {
			return 1.0
		}
	} else {
		if ind >= ind1 {
			return 2.0
		} else {
			return 3.0
		}
	}
}
