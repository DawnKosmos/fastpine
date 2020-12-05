package series

//Vix implement the Volatility Index
func Vix(close, low Series, l int) Series {
	return DivF(Sub(Highest(close, l), low), Highest(close, l), 100)
}

func Dai(open, high, low, close, volume Series, l1, l2 int) Series {
	sh := Vwma(Sum(Sub(high, close), l1), volume, l2)
	lo := Vwma(Sum(Sub(close, low), l1), volume, l2)
	return DivF(Sub(lo, sh), sh, 4)
}

func AwesomeOscillator(hl2 Series, len1, len2 int) Series {
	return Sub(Sma(hl2, len1), Sma(hl2, len2))
}
