package series

import (
	"math"

	"github.com/dawnkosmos/fastpine/exchange"
)

/*All Aritmetic Operations are from the type of the ARIT struct
This class gets feeded with a function which gets 2 float64 values and returns one
I implemented a tone but it easily can be extended by crazier kind of stuff
*/

type ARIT struct {
	src        Series
	src2       Series
	op         func(v1 float64, v2 float64) float64
	starttime  int64
	isConstant bool
	c          float64

	ug *exchange.UpdateGroup
}

func Arit(operator func(float64, float64) float64, src Series, v Value) *ARIT {
	var s ARIT
	s.src = src
	s.op = operator
	s.ug = src.UpdateGroup()

	switch v := v.(type) {
	case float64:
		s.starttime = src.Starttime()
		s.c = v
		s.isConstant = true
	case int:
		s.starttime = src.Starttime()
		s.c = float64(v)
		s.isConstant = true
	case Series:
		s.starttime = bigger(src.Starttime(), v.Starttime())
		s.src2 = v
		s.isConstant = false
	}
	return &s
}

func (s *ARIT) Starttime() int64 {
	return s.starttime
}

func (s *ARIT) Resolution() int {
	return s.src.Resolution()
}

func (s *ARIT) Value(index int) float64 {
	if s.isConstant {
		return s.op(s.src.Value(index), s.c)
	} else {
		return s.op(s.src.Value(index), s.src2.Value(index))
	}
}

func (s *ARIT) UpdateGroup() *exchange.UpdateGroup {
	return s.ug
}

func (s *ARIT) Data() *[]float64 {
	if s.isConstant {
		f := s.src.Data()
		fOut := make([]float64, 0, len(*f))
		for _, v := range *f {
			fOut = append(fOut, s.op(v, s.c))
		}
		return &fOut
	} else {
		f1 := *s.src.Data()
		f2 := *s.src2.Data()
		l1, l2 := len(f1), len(f2)
		if l1 >= l2 {
			fOut := make([]float64, l2, l2)
			f1 = f1[l1-l2:]
			for i, v := range f2 {
				fOut[i] = s.op(f1[i], v)
			}
			return &fOut
		} else {
			fOut := make([]float64, l1, l1)
			f2 = f2[l2-l1:]
			for i, v := range f1 {
				fOut[i] = s.op(v, f2[i])
			}
			return &fOut
		}
	}
}

//OPERATIONS

func Add(src Series, v Value) *ARIT {
	o := func(v1 float64, v2 float64) float64 {
		return v1 + v2
	}
	return Arit(o, src, v)
}

func Mult(src Series, v Value) *ARIT {
	o := func(v1 float64, v2 float64) float64 {
		return v1 * v2
	}
	return Arit(o, src, v)
}

func Sub(src Series, v Value) *ARIT {
	o := func(v1 float64, v2 float64) float64 {
		return v1 - v2
	}
	return Arit(o, src, v)
}

func Div(src Series, v Value) *ARIT {
	o := func(v1 float64, v2 float64) float64 {
		if v2 == 0 {
			v1 = 0
			v2 = 1
		}
		return v1 / v2
	}
	return Arit(o, src, v)
}

func Mod(src Series, v Value) *ARIT {
	o := math.Mod
	return Arit(o, src, v)
}

func Pow(src Series, v Value) *ARIT {
	o := math.Pow
	return Arit(o, src, v)
}

func Round(src Series) *ARIT {
	o := func(v1 float64, _ float64) float64 {
		return math.Round(v1)
	}
	return Arit(o, src, 0.0)
}

func Min(src Series, v Value) *ARIT {
	o := math.Min /*func(v1 float64, v2 float64) float64 {
		return math.Min(v1, v2)
	}*/
	return Arit(o, src, v)
}

func Max(src Series, v Value) *ARIT {
	o := math.Max
	return Arit(o, src, v)
}

func Remainder(src Series, v Value) *ARIT {
	o := math.Remainder
	return Arit(o, src, v)
}

func Hypot(src Series, v Value) *ARIT {
	o := math.Hypot
	return Arit(o, src, v)
}

//F functions but an extra factor to multiply the first Value
func AddF(src Series, v Value, factor float64) *ARIT {
	o := func(v1 float64, v2 float64) float64 {
		return v1*factor + v2
	}
	return Arit(o, src, v)
}

func SubF(src Series, v Value, factor float64) *ARIT {
	o := func(v1 float64, v2 float64) float64 {
		return v1*factor - v2
	}
	return Arit(o, src, v)
}

func MultF(src Series, v Value, factor float64) *ARIT {
	o := func(v1 float64, v2 float64) float64 {
		return v1 * factor * v2
	}
	return Arit(o, src, v)
}

func DivF(src Series, v Value, factor float64) *ARIT {
	o := func(v1 float64, v2 float64) float64 {
		return v1 * factor / v2
	}
	return Arit(o, src, v)
}

//C functions but an extra constant gets added to the solution
func AddC(src Series, v Value, constant float64) *ARIT {
	o := func(v1 float64, v2 float64) float64 {
		return v1 + v2 + constant
	}
	return Arit(o, src, v)
}

func SubC(src Series, v Value, constant float64) *ARIT {
	o := func(v1 float64, v2 float64) float64 {
		return v1 + v2 + constant
	}
	return Arit(o, src, v)
}

func DivC(src Series, v Value, constant float64) *ARIT {
	o := func(v1 float64, v2 float64) float64 {
		return v1 + v2 + constant
	}
	return Arit(o, src, v)
}

func MultC(src Series, v Value, constant float64) *ARIT {
	o := func(v1 float64, v2 float64) float64 {
		return v1 + v2 + constant
	}
	return Arit(o, src, v)
}
