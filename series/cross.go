package series

//Crossover equivalent to crossover(s1,s2)
func Crossover(s1 Series, v Value) Condition {
	s2, ok := v.(Series)

	if ok {
		s11, s21 := Offset(s1, 1), Offset(s2, 1)
		return And(GreaterEqual(s1, s2), Smaller(s11, s21))
	}

	s11 := Offset(s1, 1)
	return And(GreaterEqual(s1, v), Smaller(s11, v))
}

//Crossunder equivalent to crossunder(s1,s2)
func Crossunder(s1 Series, v Value) Condition {
	s2, ok := v.(Series)

	if ok {
		s11, s21 := Offset(s1, 1), Offset(s2, 1)
		return And(SmallerEqual(s1, s2), Greater(s11, s21))
	}

	s11 := Offset(s1, 1)
	return And(SmallerEqual(s1, v), Greater(s11, v))
}
