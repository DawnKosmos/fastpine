package coinbase

import (
	"fmt"
	"time"

	"github.com/dawnkosmos/fastpine/exchange"
	coinbasepro "github.com/preichenberger/go-coinbasepro/v2"
)

func (cb *COINBASE) OHCLV(ticker string, res int, start int64, end int64) ([]exchange.Candle, error) {
	if end > time.Now().Unix() {
		end = time.Now().Unix()
	}

	sss := int(end-start) / res
	var historicalPrices []exchange.Candle = make([]exchange.Candle, 0, sss)

	if end > time.Now().Unix() {
		end = time.Now().Unix()
	}

	newRes := checkResolution(res)
	var fakeEnd int64

	for start < end {
		fakeEnd = start + int64(newRes*299)
		if fakeEnd > end {
			fakeEnd = end
		}

		params := coinbasepro.GetHistoricRatesParams{
			Start:       time.Unix(start, 0),
			End:         time.Unix(fakeEnd, 0),
			Granularity: newRes,
		}

		result, err := cb.Client.GetHistoricRates(ticker, params)
		if err != nil {
			return historicalPrices, err
		}
		historicalPrices = append(historicalPrices, coinbaseCandleToCandle(result)...)
		start = start + int64(newRes*300)
		time.Sleep(time.Millisecond * 200)
	}
	correctedHP := make([]exchange.Candle, 0, sss)
	correctedHP = append(correctedHP, historicalPrices[0])
	i := 0
	for _, c := range historicalPrices[1:] {
		if c.StartTime.Unix()-correctedHP[i].StartTime.Unix() != int64(newRes) {
			kek := approximateCandles(correctedHP[i], c, newRes)
			correctedHP = append(correctedHP, kek[1])
			i++
			fmt.Println(correctedHP[i-1].StartTime, kek[1].StartTime, c.StartTime)
		}
		correctedHP = append(correctedHP, c)
		i++
	}

	kek, err := exchange.ConvertChartResolution(int64(newRes), int64(res), correctedHP)
	return kek, err
}

func avg(f1, f2 float64) float64 {
	return (f1 + f2) / 2
}

func approximateCandles(c1, c2 exchange.Candle, res int) []exchange.Candle {
	var out []exchange.Candle
	out = append(out, c1)
	for i := 0; i < int(c2.StartTime.Unix()-c1.StartTime.Unix())/res-1; i++ {
		newCandle := exchange.Candle{
			Open:      avg(out[i].Open, c2.Open),
			Close:     avg(c2.Close, out[i].Close),
			High:      avg(out[i].High, c2.High),
			Low:       avg(out[i].Low, c2.Low),
			Volume:    avg(out[i].Volume, c2.Volume),
			StartTime: time.Unix(out[i].StartTime.Unix()+int64(res), 0),
		}
		out = append(out, newCandle)
	}

	return out
}

func (b *COINBASE) Name() string {
	return "Coinbase"
}

func (b *COINBASE) Actual(Ticker string, resolution int64) (exchange.Candle, error) {
	return exchange.Candle{}, nil
}

func coinbaseCandleToCandle(ch []coinbasepro.HistoricRate) []exchange.Candle {
	newChart := make([]exchange.Candle, 0, len(ch))
	var ec exchange.Candle

	for i := len(ch) - 1; i >= 0; i-- {
		c := &ch[i]
		ec = exchange.Candle{
			Close:     c.Close,
			Open:      c.Open,
			High:      c.High,
			Low:       c.Low,
			StartTime: c.Time,
			Volume:    c.Volume,
		}
		newChart = append(newChart, ec)
	}

	return newChart
}

//{60, 300, 900, 3600, 21600, 86400}

func checkResolution(res int) int {
	var newRes int
	if res == 3600 || res == 21600 || res == 86400 || res == 300 || res == 60 || res == 900 {
		newRes = res
		return newRes
	}
	if res >= 86400 && res%86400 == 0 {
		return 86400
	}

	if res >= 21600 && res%21600 == 0 {
		return 21600
	}

	if res >= 3600 && res%3600 == 0 {
		return 3600
	}
	if res >= 300 && res%300 == 0 {
		return 300
	}
	if res >= 900 && res%900 == 0 {
		return 900
	}

	if res >= 15 && res%15 == 0 {
		return 15
	}
	return 3600
}
