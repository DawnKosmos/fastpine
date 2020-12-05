package series

/*All Comparision Operations are form the type of the comp struct
This class gets feeded with a function which Compares 2 float64 values
I just implemented ==, >=, >, <, <=, != but it easily can be extended by
crazier kind of stuff
*/

type comp struct {
	src  Series
	src2 Series
	op   func(v1 float64, v2 float64) bool

	starttime  int64
	isConstant bool
	c          float64

	USR
}

/*Comp lets you create your own logical operations. Here an example:
func Less(src Series, v Value) *comp {
	o := func(v1 float64, v2 float64) bool {
		return v1 < v2
	}
	return Comp(o, src, v)
}
@parameter gets a func(float64,float64)bool a function that takes 2 floats and returns a bool
@parameter src only accepts structs that implement the series interface
@parameter v Value accepts types of series, integer or float64
*/
func Comp(operator func(float64, float64) bool, src Series, v Value) Condition {
	var s comp
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

func (s *comp) Resolution() int {
	return s.src.Resolution()
}

func (s *comp) ValueB(index int) bool {
	if s.isConstant {
		return s.op(s.src.Value(index), s.c)
	} else {
		return s.op(s.src.Value(index), s.src2.Value(index))
	}
}

func (s *comp) DataB() []bool {
	if s.isConstant {
		f := s.src.Data()
		fOut := make([]bool, 0, len(f))
		for _, v := range f {
			fOut = append(fOut, s.op(v, s.c))
		}
		return fOut
	} else {
		f1 := s.src.Data()
		f2 := s.src2.Data()
		l1, l2 := len(f1), len(f2)
		if l1 >= l2 {
			fOut := make([]bool, l2, l2)
			f1 = f1[l1-l2:]
			for i, v := range f2 {
				fOut[i] = s.op(f1[i], v)
			}
			return fOut
		} else {
			fOut := make([]bool, l1, l1)
			f2 = f2[l2-l1:]
			for i, v := range f1 {
				fOut[i] = s.op(f2[i], v)
			}
			return fOut
		}
	}
}

//Comparison functions

//Smaller (src,v) => src < v
func Smaller(src Series, v Value) Condition {
	o := func(v1 float64, v2 float64) bool {
		return v1 < v2
	}
	return Comp(o, src, v)
}

//Greater (src,v) => src > v
func Greater(src Series, v Value) Condition {
	o := func(v1 float64, v2 float64) bool {
		return v1 > v2
	}
	return Comp(o, src, v)
}

//Equal (src,v) => src == v
func Equal(src Series, v Value) Condition {
	o := func(v1 float64, v2 float64) bool {
		return v1 == v2
	}
	return Comp(o, src, v)
}

//NotEqual (src,v) => src != v
func NotEqual(src Series, v Value) Condition {
	o := func(v1 float64, v2 float64) bool {
		return v1 != v2
	}
	return Comp(o, src, v)
}

//LessEqual (src,v) => src <= v
func SmallerEqual(src Series, v Value) Condition {
	o := func(v1 float64, v2 float64) bool {
		return v1 <= v2
	}
	return Comp(o, src, v)
}

//GreaterEqual (src,v) => src >= v
func GreaterEqual(src Series, v Value) Condition {
	o := func(v1 float64, v2 float64) bool {
		return v1 >= v2
	}
	return Comp(o, src, v)
}
