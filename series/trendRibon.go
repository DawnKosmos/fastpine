package series

/*
Source Code:
https://www.tradingview.com/script/ZsKsLiUU-noro-s-trend-ribbon-strategy/
*/

//TrendRibonNoro, The link for this script is posted in the src Code
func TrendRibonNoro(MAFunction func(src Series, len int) Series, src Series, len int) (lowerLine, upperLine Series) {
	ma := MAFunction(src, len)
	upperLine = Highest(ma, len)
	lowerLine = Lowest(ma, len)
	return
}
