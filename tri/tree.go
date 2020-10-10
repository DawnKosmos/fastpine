package tri

import (
	"fmt"
)

type Tree struct {
	L *Tree
	V Valuer
	R *Tree
}

type Valuer interface {
	Value() int
	String() string
}

func New(v Valuer) *Tree {
	var t *Tree
	insert(t, v)
	return t
}

func insert(t *Tree, v Valuer) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v.Value() < t.V.Value() {
		t.L = insert(t.L, v)
	} else {
		t.R = insert(t.R, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.L != nil {
		s += t.L.String() + " "
	}
	s += fmt.Sprint(t.V)
	if t.R != nil {
		s += " " + t.R.String()
	}
	return "(" + s + ")"
}
