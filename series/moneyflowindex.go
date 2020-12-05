package series

func MFI(src Series, volume Series, len int) Series {
	ch := Change(src, 1)
	con := SmallerEqual(ch, 0.0)
	upper := Sum(Mult(volume, Iff(con, 0, src)), len)
	lower := Sum(Mult(volume, Iff(Not(con), 0, src)), len)
	mfr := Div(upper, lower)
	mfi := DivF(mfr, Add(mfr, 1), 100)
	return mfi
}
