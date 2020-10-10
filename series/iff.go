package series

import "github.com/dawnkosmos/fastpine/exchange"

type ConCase int

const (
	CC ConCase = iota
	SC
	CS
	SS
)

type IFF struct {
	con Condition

	src1, src2 Series
	c1, c2     float64

	isConstant1, isConstant2 bool

	fOut      *[]float64
	starttime int64
	ug        *exchange.UpdateGroup

	coca ConCase
}

func Iff(con Condition, stat1 Value, stat2 Value) *IFF {
	var s IFF
	s.con = con
	s.ug = con.UpdateGroup()

	conData := *con.DataB()
	lCon := len(conData)
	fOut := make([]float64, 0, lCon)
	var f1, f2 []float64
	var l1, l2 int

	switch stat1 := stat1.(type) {
	case float64:
		s.c1 = stat1
		s.starttime = con.Starttime()
		s.isConstant1 = true
	case int:
		s.c1 = float64(stat1)
		s.isConstant1 = true
		s.starttime = con.Starttime()
	case Series:
		s.src1 = stat1
		s.isConstant1 = false
		s.starttime = bigger(con.Starttime(), stat1.Starttime())
		f1 = *stat1.Data()
		l1 = len(f1)
	}

	switch stat2 := stat2.(type) {
	case float64:
		s.c2 = stat2
		s.isConstant2 = true
	case int:
		s.c2 = float64(stat2)
		s.isConstant2 = true
	case Series:
		s.src2 = stat2
		s.isConstant2 = false
		s.starttime = bigger(con.Starttime(), stat2.Starttime())
		f2 = *stat2.Data()
		l2 = len(f2)
	}

	longest := checkLongest(lCon, l1, l2)

	if s.isConstant1 && s.isConstant2 {
		for _, v := range conData {
			if v {
				fOut = append(fOut, s.c1)
			} else {
				fOut = append(fOut, s.c2)
			}
		}
		s.coca = CC
		s.fOut = &fOut
		return &s
	}

	if !s.isConstant1 && s.isConstant2 {
		f1 = f1[longest-l1:]
		conData = conData[longest-lCon:]
		for i, v := range conData {
			if v {
				fOut = append(fOut, f1[i])
			} else {
				fOut = append(fOut, s.c2)
			}
		}
		s.coca = SC
		s.fOut = &fOut
		return &s
	}

	if s.isConstant1 && !s.isConstant2 {
		f2 = f2[longest-l2:]
		conData = conData[longest-lCon:]
		for i, v := range conData {
			if v {
				fOut = append(fOut, f2[i])
			} else {
				fOut = append(fOut, s.c1)
			}
		}
		s.coca = CS
		s.fOut = &fOut
		return &s
	}

	f1 = f1[longest-l1:]
	f2 = f2[longest-l2:]
	conData = conData[longest-lCon:]
	for i, v := range conData {
		if v {
			fOut = append(fOut, f1[i])
		} else {
			fOut = append(fOut, f2[i])
		}
		s.coca = SS
		s.fOut = &fOut

	}
	return &s
}

func (s *IFF) Value(index int) float64 {
	r := 0.0
	switch s.coca {
	case CC:
		r = s.iffresult(index, s.c1, s.c2)
	case SC:
		r = s.iffresult(index, s.src1.Value(index), s.c2)
	case CS:
		r = s.iffresult(index, s.c1, s.src2.Value(index))
	case SS:
		r = s.iffresult(index, s.src1.Value(index), s.src2.Value(index))
	}

	return r
}

func (s *IFF) iffresult(index int, a, b float64) float64 {
	if s.con.ValueB(index) {
		return a
	} else {
		return b
	}
}

func (s *IFF) Data() *[]float64 {
	return s.fOut
}

func (s *IFF) Starttime() int64 {
	return s.starttime
}

func (s *IFF) Resolution() int {
	return s.con.Resolution()
}

func checkLongest(f ...int) int {
	l := 0
	for _, v := range f {
		if v > l {
			l = v
		}
	}
	return l
}
