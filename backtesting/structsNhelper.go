package backtesting

type Candle struct {
	Timestamp int
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Volume    float64
	Indicator []float64
	prev      *Candle
	next      *Candle
}

//Helper Functions

func lowestLen(in ...[]float64) int {
	l := len(in[0])
	for _, v := range in {
		if len(v) < l {
			l = len(v)
		}
	}
	return l
}
