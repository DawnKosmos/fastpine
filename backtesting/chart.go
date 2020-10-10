package backtesting

type Chart struct {
	exchange        string
	ticker          string
	Resolution      int
	Starttime       int
	IndicatorLayout map[string]int
	ch              []Candle
}

type Candle struct {
	Timestamp int
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Indicator []float64
	prev      *Candle
	next      *Candle
}
