package backtesting

import (
	"github.com/dawnkosmos/fastpine/exchange"
	"github.com/dawnkosmos/fastpine/series"
)

type Chart struct {
	Exchange   string
	Ticker     string
	Resolution int
	Starttime  int
	src        []*Candle
}

func exchangeCandleToCandle(ch []exchange.Candle, indi ...series.Series) []*Candle {
	var indicator [][]float64
	for _, v := range indi {
		indicator = append(indicator, v.Data())
	}
	l := lowestLen(indicator...)
	var out []*Candle = make([]*Candle, 0, l+1)

	ch = ch[len(ch)-l:]
	for _, v := range indicator {
		v = v[len(v)-l:]
	}

	out = append(out, &Candle{})

	for i, v := range ch {
		var c Candle
		c.Open, c.Close, c.High, c.Low, c.Volume = v.Open, v.Close, v.High, v.Low, v.Volume
		c.Timestamp = int(v.StartTime.Unix())
		var f []float64 = make([]float64, 0, len(indicator))
		for _, j := range indicator {
			f = append(f, j[i])
		}
		c.Indicator = f

		c.prev = out[i]
		out[i].next = &c
		out = append(out, &c)
	}
	return out
}
