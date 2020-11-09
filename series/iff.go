package series

import (
	"github.com/dawnkosmos/fastpine/helper"
)

type IffCase int

const (
	CONSTANTCONSTANT IffCase = iota
	SERIESCONSTANT
	CONSTANTSERIES
	SERIESSERIES
	CONDITIONBOOL
	BOOLCONDITION
	CONDITIONCONDITION
)

type IFF struct {
	con Condition

	src1, src2   Series
	bsrc1, bsrc2 Condition

	c1, c2 float64
	b1, b2 bool

	isConstant1, isConstant2 bool
	isBool1, isBool2         bool
	fOut                     []float64
	bOut                     []bool

	USR

	coca IffCase
}

/*Iff is equivalent to
if con {
	stat1
} else {
	stat2
}

Iff implements the Condition and Series interface. It can have some buggs atm
*/
func Iff(con Condition, stat1 Value, stat2 Value) *IFF {
	var s IFF
	s.con = con
	s.ug = con.UpdateGroup()

	conData := con.DataB()
	lCon := len(conData)
	var fOut []float64
	var bOut []bool
	var f1, f2 []float64
	var b1, b2 []bool

	switch stat1 := stat1.(type) {
	case float64:
		s.c1 = stat1
		s.isConstant1 = true
		s.starttime = con.Starttime()
	case int:
		s.c1 = float64(stat1)
		s.isConstant1 = true
		s.starttime = con.Starttime()
	case Series:
		s.src1 = stat1
		s.starttime = helper.Int64Max(con.Starttime(), stat1.Starttime())
		f1 = stat1.Data()
	case Condition:
		s.isBool1 = true
		s.bsrc1 = stat1
		s.starttime = helper.Int64Max(con.Starttime(), stat1.Starttime())
		b1 = stat1.DataB()
	case bool:
		s.b1 = stat1
		s.starttime = con.Starttime()
		s.isBool1 = true
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
		s.starttime = helper.Int64Max(s.starttime, stat2.Starttime())
		f2 = stat2.Data()
	case Condition:
		s.bsrc2 = stat2
		s.isConstant2 = false
		s.starttime = helper.Int64Max(s.starttime, stat2.Starttime())
		b2 = stat2.DataB()
	case bool:
		s.b2 = stat2
		s.isBool2 = true
	}

	if s.isBool1 && len(b2) > 1 {
		bOut = make([]bool, 0, len(conData))
		shortest := helper.Min(len(b2), len(conData))

		conData = conData[len(conData)-shortest:]
		b2 = b2[len(b2)-shortest:]

		for i, v := range conData {
			if v {
				bOut = append(bOut, s.b1)
			} else {
				bOut = append(bOut, b2[i])
			}
		}
		s.coca = BOOLCONDITION
		s.bOut = bOut
		return &s
	}

	if s.isBool2 && len(b1) > 1 {
		shortest := helper.Min(len(b1), len(conData))
		conData = conData[len(conData)-shortest:]
		b1 = b1[len(b1)-shortest:]

		bOut = make([]bool, 0, len(conData))
		for i, v := range conData {
			if v {
				bOut = append(bOut, b1[i])
			} else {
				bOut = append(bOut, s.b2)
			}
		}
		s.coca = CONDITIONBOOL
		s.bOut = bOut
		return &s
	}

	if len(b2) > 1 && len(b1) > 1 {
		shortest := helper.Min(len(b1), len(conData), len(b2))
		conData = conData[len(conData)-shortest:]
		b1 = b1[len(b1)-shortest:]
		b2 = b2[len(b2)-shortest:]

		bOut = make([]bool, 0, len(conData))
		for i, v := range conData {
			if v {
				bOut = append(bOut, b1[i])
			} else {
				bOut = append(bOut, b2[i])
			}
		}
		s.coca = CONDITIONCONDITION
		s.bOut = bOut
		return &s
	}

	if s.isConstant1 && s.isConstant2 {
		fOut = make([]float64, 0, len(conData))
		for _, v := range conData {
			if v {
				fOut = append(fOut, s.c1)
			} else {
				fOut = append(fOut, s.c2)
			}
		}
		s.coca = CONSTANTCONSTANT
		s.fOut = fOut
		return &s
	}

	if !s.isConstant1 && s.isConstant2 {
		fOut = make([]float64, 0, len(conData))

		shortest := helper.Min(len(f1), len(conData))
		conData = conData[len(conData)-shortest:]
		f1 = f1[len(f1)-shortest:]

		for i, v := range conData {
			if v {
				fOut = append(fOut, f1[i])
			} else {
				fOut = append(fOut, s.c2)
			}
		}
		s.coca = SERIESCONSTANT
		s.fOut = fOut
		return &s
	}

	if s.isConstant1 && !s.isConstant2 {

		shortest := helper.Min(len(conData), len(f2))
		conData = conData[len(conData)-shortest:]
		f2 = f2[len(f2)-shortest:]

		conData = conData[lCon-shortest:]
		for i, v := range conData {
			if v {
				fOut = append(fOut, f2[i])
			} else {
				fOut = append(fOut, s.c1)
			}
		}
		s.coca = CONSTANTSERIES
		s.fOut = fOut
		return &s
	}

	for i, v := range conData {

		shortest := helper.Min(len(f1), len(conData), len(f2))
		conData = conData[len(conData)-shortest:]
		f1 = f1[len(f1)-shortest:]
		f2 = f2[len(f2)-shortest:]

		if v {
			fOut = append(fOut, f1[i])
		} else {
			fOut = append(fOut, f2[i])
		}
		s.coca = SERIESSERIES
		s.fOut = fOut

	}
	return &s
}

func (s *IFF) Value(index int) float64 {
	r := 0.0
	switch s.coca {
	case CONSTANTCONSTANT:
		r = s.IFFresult(index, s.c1, s.c2)
	case SERIESCONSTANT:
		r = s.IFFresult(index, s.src1.Value(index), s.c2)
	case CONSTANTSERIES:
		r = s.IFFresult(index, s.c1, s.src2.Value(index))
	case SERIESSERIES:
		r = s.IFFresult(index, s.src1.Value(index), s.src2.Value(index))
	}

	return r
}

func (s *IFF) ValueB(index int) bool {
	var r bool
	switch s.coca {
	case CONDITIONBOOL:
		r = s.BIffresult(index, s.bsrc1.ValueB(index), s.b2)
	case BOOLCONDITION:
		r = s.BIffresult(index, s.b1, s.bsrc2.ValueB(index))
	case CONDITIONCONDITION:
		r = s.BIffresult(index, s.bsrc1.ValueB(index), s.bsrc2.ValueB(index))
	}

	return r
}

func (s *IFF) DataB() []bool {
	return s.bOut
}

func (s *IFF) IFFresult(index int, a, b float64) float64 {
	if s.con.ValueB(index) {
		return a
	} else {
		return b
	}
}

func (s *IFF) BIffresult(index int, a, b bool) bool {
	if s.con.ValueB(index) {
		return a
	} else {
		return b
	}
}

func (s *IFF) Data() []float64 {
	return s.fOut
}

func (s *IFF) Resolution() int {
	return s.con.Resolution()
}

func lowestInt(f ...int) int {
	l := f[0]
	for _, v := range f[1:] {
		if v < l {
			l = v
		}
	}
	return l
}
