package series

//DoubleMA returns the double smoothed version of an MA
func DoubleMA(op func(Series, int) Series, src Series, l int) Series {
	e1 := op(src, l)
	return SubF(e1, op(e1, l), 2)
}
