package series

import (
	"math"

	"github.com/dawnkosmos/fastpine/exchange"
)

/*All Aritmetic Operations are from the type of the arit struct
This class gets feeded with a function which gets 2 float64 values and returns one
I implemented a tone but it easily can be extended by crazier kind of stuff
*/

type arit struct {
	src        Series
	src2       Series
	op         func(v1 float64, v2 float64) float64
	starttime  int64
	isConstant bool
	c          float64

	ug *exchange.UpdateGroup
}

/*Arit lets you create your own Arithmetic operations. Here an example:

@parameter v Value accepts types of series, integer or float64

func Add(src Series, v Value) Series {
	o := func(v1 float64, v2 float64) float64 {
		return v1 + v2
	}
	return Arit(o, src, v)
}
*/
func Arit(operator func(float64, float64) float64, src Series, v Value) Series {
	var s arit
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

func (s *arit) Starttime() int64 {
	return s.starttime
}

func (s *arit) Resolution() int {
	return s.src.Resolution()
}

func (s *arit) Value(index int) float64 {
	if s.isConstant {
		return s.op(s.src.Value(index), s.c)
	} else {
		return s.op(s.src.Value(index), s.src2.Value(index))
	}
}

func (s *arit) UpdateGroup() *exchange.UpdateGroup {
	return s.ug
}

func (s *arit) Data() []float64 {
	if s.isConstant {
		f := s.src.Data()
		fOut := make([]float64, 0, len(f))
		for _, v := range f {
			fOut = append(fOut, s.op(v, s.c))
		}
		return fOut
	} else {
		f1 := s.src.Data()
		f2 := s.src2.Data()
		l1, l2 := len(f1), len(f2)
		if l1 >= l2 {
			fOut := make([]float64, l2, l2)
			f1 = f1[l1-l2:]
			for i, v := range f2 {
				fOut[i] = s.op(f1[i], v)
			}
			return fOut
		} else {
			fOut := make([]float64, l1, l1)
			f2 = f2[l2-l1:]
			for i, v := range f1 {
				fOut[i] = s.op(v, f2[i])
			}
			return fOut
		}
	}
}

//OPERATIONS

func Abs(src Series) Series {
	o := func(v1 float64, _ float64) float64 {
		return math.Abs(v1)
	}
	return Arit(o, src, 0)
}

//Add (a,b) => a + b
func Add(src Series, v Value) Series {
	o := func(v1 float64, v2 float64) float64 {
		return v1 + v2
	}
	return Arit(o, src, v)
}

//Mult (a,b) => a*b
func Mult(src Series, v Value) Series {
	o := func(v1 float64, v2 float64) float64 {
		return v1 * v2
	}
	return Arit(o, src, v)
}

//Sub (src,v) => src - v
func Sub(src Value, v Value) Series {
	s, ok := src.(Series)

	if ok {
		o := func(v1 float64, v2 float64) float64 {
			return v1 - v2
		}
		return Arit(o, s, v)
	}

	f, ok := src.(float64)

	if ok {
		v, ok := src.(Series)
		if ok {
			o := func(v1 float64, v2 float64) float64 {
				return v2 - v1
			}
			return Arit(o, v, f)
		}
	}

	kek, ok := src.(int)
	if ok {
		v, ok := src.(Series)
		if ok {
			o := func(v1 float64, v2 float64) float64 {
				return v2 - v1
			}
			return Arit(o, v, kek)
		}
	}

	return nil

}

//Div (src,v) => src/v
func Div(src Series, v Value) Series {
	o := func(v1 float64, v2 float64) float64 {
		if v2 == 0 {
			v1 = 0
			v2 = 1
		}
		return v1 / v2
	}
	return Arit(o, src, v)
}

//Mod (src,v) => src%v
func Mod(src Series, v Value) Series {
	o := math.Mod
	return Arit(o, src, v)
}

//Pow (src,v) => src^v
func Pow(src Series, v Value) Series {
	o := math.Pow
	return Arit(o, src, v)
}

//Round (src) => round(src)
func Round(src Series) Series {
	o := func(v1 float64, _ float64) float64 {
		return math.Round(v1)
	}
	return Arit(o, src, 0.0)
}

//Min (src,v) => min(src,v)
func Min(src Series, v Value) Series {
	o := math.Min /*func(v1 float64, v2 float64) float64 {
		return math.Min(v1, v2)
	}*/
	return Arit(o, src, v)
}

//Max (src,v) => max(src,v)
func Max(src Series, v Value) Series {
	o := math.Max
	return Arit(o, src, v)
}

//Remainder (src,v) => return remainder of src/v
func Remainder(src Series, v Value) Series {
	o := math.Remainder
	return Arit(o, src, v)
}

//Hypot (src, v) => sqrt(src²+v²)
func Hypot(src Series, v Value) Series {
	o := math.Hypot
	return Arit(o, src, v)
}

//AddF (src, v, factor) => src*factor + v
func AddF(src Series, v Value, factor float64) Series {
	o := func(v1 float64, v2 float64) float64 {
		return v1*factor + v2
	}
	return Arit(o, src, v)
}

//SubF (src,v,factor) => src*factor - v
func SubF(src Series, v Value, factor float64) Series {
	o := func(v1 float64, v2 float64) float64 {
		return v1*factor - v2
	}
	return Arit(o, src, v)
}

//MultF (src,v,factor) => src*factor * v
func MultF(src Series, v Value, factor float64) Series {
	o := func(v1 float64, v2 float64) float64 {
		return v1 * factor * v2
	}
	return Arit(o, src, v)
}

//DivF (src,v,factor) => src*factor/v
func DivF(src Series, v Value, factor float64) Series {
	o := func(v1 float64, v2 float64) float64 {
		return v1 * factor / v2
	}
	return Arit(o, src, v)
}

//AddC (src,v,constant) => src+v+constant
func AddC(src Series, v Value, constant float64) Series {
	o := func(v1 float64, v2 float64) float64 {
		return v1 + v2 + constant
	}
	return Arit(o, src, v)
}

//SubC (src,v,constant) => src-v+constant
func SubC(src Series, v Value, constant float64) Series {
	o := func(v1 float64, v2 float64) float64 {
		return v1 - v2 + constant
	}
	return Arit(o, src, v)
}

//DivC (src,v,constant) => src/v+constant
func DivC(src Series, v Value, constant float64) Series {
	o := func(v1 float64, v2 float64) float64 {
		return v1/v2 + constant
	}
	return Arit(o, src, v)
}

//MultC (src,v,constant) => src*v+constant
func MultC(src Series, v Value, constant float64) Series {
	o := func(v1 float64, v2 float64) float64 {
		return v1*v2 + constant
	}
	return Arit(o, src, v)
}

//Sqrt (src) => sqrt(src)
func Sqrt(src Series) Series {
	o := func(v1, v2 float64) float64 {
		return math.Sqrt(v1)
	}
	return Arit(o, src, 0)
}

func Neg(src Series) Series {
	o := func(v1, v2 float64) float64 {
		return -v1
	}
	return Arit(o, src, 0)
}
