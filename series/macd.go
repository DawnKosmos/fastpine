package series

//Macd is the equivalent of macd(source, fastlen, slowlen, siglen). Returns the macd, signal & histogram value
func Macd(src Series, fastLen, slowLen, SignalLen int) (macd Series, signal Series, histogram Series) {
	f := Ema(src, fastLen)
	s := Ema(src, slowLen)
	// macd = f - s
	macd = Sub(f, s)
	signal = Ema(macd, SignalLen)
	//histogram macd - signal
	histogram = Sub(macd, signal)
	return
}
