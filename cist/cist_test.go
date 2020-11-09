package cist

import (
	"fmt"
	"testing"
)

func TestCist(t *testing.T) {
	l := New()
	a := []float64{1.0, 1.1, 1.2, 1.3}
	s := []float64{2.0, 9.1, 2.2, 2.3}
	d := []float64{3.0, 3.1, 3.2, 3.3}

	l.InitData([]float64{1, 2, 3, 4, 5, 6})
	l.FillElements(a, s, d)
	for i := 0; i < 4; i++ {
		fmt.Println(i, l.GetEle(i).Value())
	}

	l.Update(6.0, 1.0, 1.0, 1.0, 1.0)
	l.Add()
	for i := 0; i < 4; i++ {
		fmt.Println(i, l.GetEle(i).Value())
	}
	l.Update(7.0, 2.0, 1.0, 1.0, 1.0)
	l.Add()
	l.Update(8.0, 3.0, 1.0, 1.0, 1.0)
	l.Add()
	l.Update(9.0, 4.0, 1.0, 1.0, 1.0)
	l.Add()

	for i := 0; i < 4; i++ {
		fmt.Println(i, l.GetEle(i).Value())
	}
}
