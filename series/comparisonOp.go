package series

/*All Comparision Operations are form the type of the COMP struct
This class gets feeded with a function which compares 2 float64 values
I just implemented ==, >=, >, <, <=, != but it easily can be extended by
crazier kind of stuff
*/

type COMP struct {
	src  Series
	src2 Series
	op   func(v1 float64, v2 float64) bool

	starttime  int64
	isConstant bool
	c          float64
}

func comp(operator func(float64, float64) bool, src Series, v Value) *COMP {
	var s COMP
	s.src = src
	s.op = operator

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

func (s *COMP) Starttime() int64 {
	return s.starttime
}

func (s *COMP) Resolution() int {
	return s.src.Resolution()
}

func (s *COMP) ValueB(index int) bool {
	if s.isConstant {
		return s.op(s.src.Value(index), s.c)
	} else {
		return s.op(s.src.Value(index), s.src2.Value(index))
	}
}

func (s *COMP) DataB() *[]bool {
	if s.isConstant {
		f := s.src.Data()
		fOut := make([]bool, 0, len(*f))
		for _, v := range *f {
			fOut = append(fOut, s.op(v, s.c))
		}
		return &fOut
	} else {
		f1 := *s.src.Data()
		f2 := *s.src2.Data()
		l1, l2 := len(f1), len(f2)
		if l1 >= l2 {
			fOut := make([]bool, l2, l2)
			f1 = f1[l1-l2:]
			for i, v := range f2 {
				fOut[i] = s.op(v, f1[i])
			}
			return &fOut
		} else {
			fOut := make([]bool, l1, l1)
			f2 = f2[l2-l1:]
			for i, v := range f1 {
				fOut[i] = s.op(v, f2[i])
			}
			return &fOut
		}
	}
}

//Comparison functions

func Less(src Series, v Value) *COMP {
	o := func(v1 float64, v2 float64) bool {
		return v1 < v2
	}
	return comp(o, src, v)
}

func Greater(src Series, v Value) *COMP {
	o := func(v1 float64, v2 float64) bool {
		return v1 > v2
	}
	return comp(o, src, v)
}
func Equal(src Series, v Value) *COMP {
	o := func(v1 float64, v2 float64) bool {
		return v1 == v2
	}
	return comp(o, src, v)
}

func NotEqual(src Series, v Value) *COMP {
	o := func(v1 float64, v2 float64) bool {
		return v1 != v2
	}
	return comp(o, src, v)
}

func LessEqual(src Series, v Value) *COMP {
	o := func(v1 float64, v2 float64) bool {
		return v1 <= v2
	}
	return comp(o, src, v)
}

func GreaterEqual(src Series, v Value) *COMP {
	o := func(v1 float64, v2 float64) bool {
		return v1 >= v2
	}
	return comp(o, src, v)
}
