package series

//Stoch is equivalent to stoch(close,high,low, len)
func Stoch(close, high, low Series, l int) Series {
	lo := Lowest(low, l)
	hi := Highest(high, l)
	return DivF(Sub(close, lo), Sub(hi, lo), 100)
}
