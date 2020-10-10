package backtesting

import (
	"math"

	"github.com/dawnkosmos/fastpine/exchange"
)

type Trade struct {
	Side        bool // true = Long, false = Short; gets translated when outputing
	EntryTime   int64
	ExitTime    int64
	EntryPrice  float64
	ExitPrice   float64
	EntryCandle exchange.Candle
	ExitCandle  exchange.Candle
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

func getTradesFromBuySell(v []exchange.Candle, buy, sell []bool) []Trade {
	v = v[len(v)-len(buy):]
	for i := 0; i < len(buy); i++ {
		if buy[i] {

		}
	}
}

func longT(v []exchange.Candle, sell []bool) Trade {
	var t Trade
	t.Side = true
	t.EntryPrice = v[0].Open
	for i := 0; i < len(sell); i++ {
		if sell {
			t.ExitCandle = v[i]
			t.ExitPrice = v[i].Close
			t.ExitTime = v[i].StartTime.Unix() + v[i]
		}
	}
}
