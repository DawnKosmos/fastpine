package series

/*
ATR should get an SMA/EMA/RMA/VWMA(has to be wrapped first) or any other MMA indicator you can imagine
*/

//Atr gets an MA function, and a TR and len
func Atr(ma func(Series, int) Series, tr *TR, l int) Series {
	return ma(tr, l)
}
