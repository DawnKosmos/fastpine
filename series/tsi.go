package series

//Tsi is the true strenght indicator
func Tsi(src Series, r, s int) Series {
	src1 := Offset(src, 1)
	m := Sub(src, src1)

	t1 := Ema(Ema(m, r), s)
	t2 := Ema(Ema(Abs(m), r), s)
	return DivF(t1, t2, 100)
}
