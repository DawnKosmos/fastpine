package deribit

import (
	"strconv"
	"time"

	"github.com/dawnkosmos/fastpine/exchange"
	"github.com/frankrap/deribit-api/models"
)

func (d *DERIBIT) OHCLV(ticker string, resolution int, start, end int64) ([]exchange.Candle, error) {
	if end > time.Now().Unix() {
		end = time.Now().Unix()
	}

	var historicalPrices []exchange.Candle = make([]exchange.Candle, 0, int(end-start)/resolution)
	res := resolution / 60
	newRes := checkResolution(res)
	var fakeEnd int64

	for start < end {
		fakeEnd = start + int64(60*newRes*9999)
		if fakeEnd > end {
			fakeEnd = end
		}

		params := &models.GetTradingviewChartDataParams{
			InstrumentName: ticker,
			StartTimestamp: start * 1000,
			EndTimestamp:   fakeEnd * 1000,
			Resolution:     resolutionToString(newRes),
		}
		result, err := d.d.GetTradingviewChartData(params)
		if err != nil {
			return historicalPrices, err
		}

		historicalPrices = append(historicalPrices, deribitCandleToCandle(result)...)
		start = start + int64(60*newRes*10000)
	}

	kek, err := exchange.ConvertChartResolution(int64(newRes), int64(res), historicalPrices)
	return kek, err
}

func (b *DERIBIT) Name() string {
	return "Deribit"
}

func (b *DERIBIT) Actual(Ticker string, resolution int64) (exchange.Candle, error) {
	return exchange.Candle{}, nil
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

func deribitCandleToCandle(c models.GetTradingviewChartDataResponse) []exchange.Candle {
	newChart := make([]exchange.Candle, 0, len(c.Ticks))
	var ec exchange.Candle
	for i := 0; i < len(c.Ticks); i++ {
		ec = exchange.Candle{
			Close:     c.Close[i],
			Open:      c.Open[i],
			High:      c.High[i],
			Low:       c.Low[i],
			StartTime: time.Unix(c.Ticks[i]/1000, 0),
			Volume:    c.Volume[i],
		}
		newChart = append(newChart, ec)
	}
	return newChart
}

//1 3 5 10 15 30 60 120 180 360 720 1D
func checkResolution(res int) int {
	var newRes int
	if res == 1 || res == 3 || res == 5 || res == 15 || res == 30 || res == 60 || res == 120 || res == 180 || res == 360 || res == 720 || res == 1440 {
		newRes = res
		return newRes
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
	if res >= 180 && res%180 == 0 {
		return 180
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

/*
type GetTradingviewChartDataParams struct {
	InstrumentName string `json:"instrument_name"`
	StartTimestamp int64  `json:"start_timestamp"`
	EndTimestamp   int64  `json:"end_timestamp"`
	Resolution     string `json:"resolution"`
}
*/
