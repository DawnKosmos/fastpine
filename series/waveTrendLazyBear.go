package series

/*
made by @LazyBear
Link:
https://www.tradingview.com/script/2KE8wTuF-Indicator-WaveTrend-Oscillator-WT/
*/
//WaveTrend by Lazybear, Check link in the source  code, default for src is HCL3
func WaveTrend(src Series, channelLen, avgLen int) (wt1, wt2 Series) {
	esa := Ema(src, channelLen)
	d := Ema(Abs(Sub(src, esa)), channelLen)
	ci := Div(Sub(src, esa), Mult(d, 0.015))
	wt1 = Ema(ci, avgLen)
	wt2 = Sma(wt1, 4)
	return
}
