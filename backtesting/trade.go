package backtesting

import (
	"math"

	"github.com/dawnkosmos/fastpine/series/strategy"
)

type Trades []*Trade

type Trade struct {
	/*Side       series.Side
	EntryPrice float64
	ExitPrice  float64

	EntryTime      int
	ExitTime       int*/
	strategy.Trade
	EntryCondition *Candle
	ExitCondition  *Candle
}

func seriesTradeToTrade(o Chart, tr []strategy.Trade) []*Trade {
	ch := o.src
	t1 := int64(ch[0].Timestamp)
	res := int64(o.Resolution)
	trades := make([]*Trade, 0, len(tr))
	for _, v := range tr {
		t := Trade{Trade: v}
		if t.ExitTime.Unix() < int64(t1) {
			continue
		}
		t.EntryCondition = ch[(t.EntryTime.Unix()-t1)%res-1]
		t.ExitCondition = ch[(t.ExitTime.Unix()-t1)%res-1]
	}
	return trades
}

func (t *Trade) GetGains() float64 {
	var x float64
	if t.Side {
		x = (t.ExitPrice - t.EntryPrice) / t.EntryPrice
	} else {
		x = -1 * (t.ExitPrice - t.EntryPrice) / t.EntryPrice
	}
	return math.Round(x*1000) / 10
}
