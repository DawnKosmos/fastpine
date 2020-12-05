package series

func BollingerBands(src Series, len int, mult float64) (LowerBand, Basis, UpperBand Series) {
	Basis = Sma(src, len)
	std := Mult(Stdev(src, len), mult)
	LowerBand = Sub(Basis, std)
	UpperBand = Add(Basis, std)
	return
}

func BollingerBandsWidth(src Series, len int, mult float64) Series {
	l, b, u := BollingerBands(src, len, mult)
	return Div(Sub(u, l), b)
}
