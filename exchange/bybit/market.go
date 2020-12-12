package bybit

import (
	"log"
	"strconv"
	"time"

	"github.com/dawnkosmos/fastpine/exchange"
	"github.com/frankrap/bybit-api/rest"
)

func (b *BYBIT) Name() string {
	return "Bybit"
}

func (b *BYBIT) Actual(Ticker string, resolution int64) (exchange.Candle, error) {
	return exchange.Candle{}, nil
}
func (b *BYBIT) OHCLV(ticker string, resolution int, start int64, end int64) ([]exchange.Candle, error) {
	if end > time.Now().Unix() {
		end = time.Now().Unix()
	}

	var historicalPrices []exchange.Candle = make([]exchange.Candle, 0, int(end-start)/resolution)
	res := resolution / 60
	newRes := checkResolution(res)

	for start < end {
		c, err := b.b.GetKLine(ticker, resolutionToString(newRes), start, 200)
		if err != nil {
			log.Printf("Error OHCLV FTX %v", err)
			return historicalPrices, err
		}
		historicalPrices = append(historicalPrices, bbCandleToCandle(c)...)
		start = start + int64(60*newRes*200)
	}

	kek, err := exchange.ConvertChartResolution(int64(newRes), int64(res), historicalPrices)
	return kek, err
}

func resolutionToString(i int) string {
	switch i {
	case 1440:
		return "D"
	case 10080:
		return "W"
	default:
		return strconv.Itoa(i)
	}
}

func deribitCandleToCandle(ch []rest.OHLC) []exchange.Candle {
	newChart := make([]exchange.Candle, 0, len(ch))
	var ec exchange.Candle
	for _, c := range ch {
		ec = exchange.Candle{
			Close:     c.Close,
			Open:      c.Open,
			High:      c.High,
			Low:       c.Low,
			StartTime: time.Unix(c.OpenTime/1000, 0),
			Volume:    c.Volume,
		}
		newChart = append(newChart, ec)
	}
	return newChart
}

func checkResolution(res int) int {
	var newRes int
	if res == 1 || res == 3 || res == 5 || res == 15 || res == 30 || res == 60 || res == 120 || res == 240 || res == 360 || res == 720 || res == 1440 || res == 10080 {
		newRes = res
		return newRes
	}
	if res >= 10080 && res%10080 == 0 {
		return 86400
	}

	if res >= 1440 && res%1440 == 0 {
		return 1440
	}

	if res >= 720 && res%720 == 0 {
		return 720
	}
	if res >= 360 && res%360 == 0 {
		return 360
	}
	if res >= 240 && res%240 == 0 {
		return 240
	}

	if res >= 120 && res%120 == 0 {
		return 120
	}
	if res >= 60 && res%60 == 0 {
		return 60
	}
	if res >= 30 && res%30 == 0 {
		return 30
	}
	if res >= 15 && res%15 == 0 {
		return 15
	}

	if res >= 5 && res%5 == 0 {
		return 5
	}
	return 60
}
