package series

import (
	"github.com/dawnkosmos/fastpine/helper"
)

type and struct {
	con1 Condition
	con2 Condition
	USR
	fOut []bool
}

/*And(con1,con2)
The equivalent of "con1 and con2"
*/
func And(con1 Condition, con2 Condition) Condition {
	var s and
	s.con1 = con1
	s.con2 = con2
	s.starttime = bigger(con1.Starttime(), con2.Starttime())
	d1, d2 := con1.DataB(), con2.DataB()
	l1, l2 := len(d1), len(d2)
	s.resolution = con1.Resolution()
	s.ug = con1.UpdateGroup()
	shortestLen := lowestInt(l1, l2)

	d1 = d1[l1-shortestLen:]
	d2 = d2[l2-shortestLen:]
	s.fOut = make([]bool, 0, shortestLen)
	for i, v := range d1 {
		s.fOut = append(s.fOut, v && d2[i])
	}

	return &s
}

func (s *and) ValueB(i int) bool {
	return s.con1.ValueB(i) && s.con2.ValueB(i)
}

func (s *and) DataB() []bool {
	return s.fOut
}

//or getting 2 conditions returning a Condition
type or struct {
	con1 Condition
	con2 Condition
	USR
	fOut []bool
}

/*Or(con1,con2)
The equivalent of "con1 or con2"
*/
func Or(con1 Condition, con2 Condition) Condition {
	var s or
	s.con1 = con1
	s.con2 = con2
	s.starttime = bigger(con1.Starttime(), con2.Starttime())
	d1, d2 := con1.DataB(), con2.DataB()
	l1, l2 := len(d1), len(d2)
	s.resolution = con1.Resolution()
	s.ug = con1.UpdateGroup()
	shortestLen := lowestInt(l1, l2)

	d1 = d1[l1-shortestLen:]
	d2 = d2[l2-shortestLen:]
	s.fOut = make([]bool, 0, shortestLen)
	for i, v := range d1 {
		s.fOut = append(s.fOut, v || d2[i])
	}

	return &s
}

func (s *or) ValueB(i int) bool {
	return s.con1.ValueB(i) || s.con2.ValueB(i)
}

func (s *or) DataB() []bool {
	return s.fOut
}

type xor struct {
	con1 Condition
	con2 Condition
	USR
	fOut []bool
}

/*Or(con1,con2)
There is no XOR in pinescript but an equivalent is (con1 or con2) and not (con1 and con2)
*/
func Xor(con1 Condition, con2 Condition) Condition {
	var s xor
	s.con1 = con1
	s.con2 = con2
	s.starttime = bigger(con1.Starttime(), con2.Starttime())
	d1, d2 := con1.DataB(), con2.DataB()
	l1, l2 := len(d1), len(d2)
	s.resolution = con1.Resolution()
	s.ug = con1.UpdateGroup()
	shortestLen := lowestInt(l1, l2)

	d1 = d1[l1-shortestLen:]
	d2 = d2[l2-shortestLen:]
	s.fOut = make([]bool, 0, shortestLen)
	for i, v := range d1 {
		s.fOut = append(s.fOut, helper.Xor(v, d2[i]))
	}

	return &s
}

func (s *xor) ValueB(i int) bool {
	return helper.Xor(s.con1.ValueB(i), s.con2.ValueB(i))
}

func (s *xor) DataB() []bool {
	return s.fOut
}

type not struct {
	con Condition
	USR
	fOut []bool
}

/*Not(con1)
The equivalent of "not con1"
*/
func Not(con Condition) Condition {
	var s not
	s.con = con
	s.starttime = con.Starttime()
	s.resolution = con.Resolution()
	s.ug = con.UpdateGroup()

	s.fOut = make([]bool, 0, len(s.con.DataB())+8)
	for _, v := range con.DataB() {
		s.fOut = append(s.fOut, !v)
	}

	return &s
}

func (s *not) ValueB(i int) bool {
	return !s.con.ValueB(i)
}

func (s *not) DataB() []bool {
	return s.fOut
}
