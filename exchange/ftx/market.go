package ftx

import (
	"log"
	"strconv"
	"time"

	"github.com/dawnkosmos/fastpine/exchange"
)

type HistoricalPrices struct {
	Success bool              `json:"success"`
	Result  []exchange.Candle `json:"result"`
}

func (client *FTX) Actual(ticker string, res int64) (exchange.Candle, error) {
	var h HistoricalPrices
	resp, err := client._get(
		"markets/"+ticker+
			"/candles?resolution="+strconv.FormatInt(res, 10)+
			"&limit=1",
		[]byte(""))
	if err != nil {
		log.Printf("Error ActualCandle FTX %v", err)
		return exchange.Candle{0, 0, 0, 0, 0, time.Now()}, err
	}
	err = _processResponse(resp, &h)
	return h.Result[0], err
}

// Resolution is sec length, window length in seconds. 15, 60(1m), 300(5m), 900(15m), 3600(60m), 14400(4h), 86400(1D)
// max amount you can ask for is 5000 candles

func (client *FTX) OHCLV(ticker string, res int, startTime int64, endTime int64) ([]exchange.Candle, error) {
	var historicalPrices []exchange.Candle
	var end int64 = 0
	newRes := checkResolution(res)

	for startTime < endTime {
		end = startTime + int64(newRes*1500)
		if end >= endTime {
			c, err := client.getHistoricalPrices(ticker, int64(newRes), startTime, endTime)
			if err != nil {
				log.Printf("Error OHCLV FTX %v", err)
				return historicalPrices, err
			}
			historicalPrices = append(historicalPrices, c...)
		} else {
			c, err := client.getHistoricalPrices(ticker, int64(newRes), startTime, end)
			if err != nil {
				log.Printf("Error OHCLV FTX %v", err)
				return historicalPrices, err
			}
			historicalPrices = append(historicalPrices, c...)

		}
		startTime = startTime + int64(newRes*1501)
	}

	kek, err := exchange.ConvertChartResolution(int64(newRes), int64(res), historicalPrices)
	return kek, err
}

func (client *FTX) getHistoricalPrices(ticker string, res int64, startTime int64, endTime int64) ([]exchange.Candle, error) {
	var historicalPrices HistoricalPrices
	resp, err := client._get(
		"markets/"+ticker+
			"/candles?resolution="+strconv.FormatInt(res, 10)+
			"&start_time="+strconv.FormatInt(startTime, 10)+
			"&end_time="+strconv.FormatInt(endTime, 10),
		[]byte(""))
	if err != nil {
		log.Printf("Error OHCLV FTX %v", err)
		return historicalPrices.Result, err
	}
	err = _processResponse(resp, &historicalPrices)
	return historicalPrices.Result, nil
}

//checkResolution looking if the asked resolution is a valid one
func checkResolution(res int) int {
	var newRes int
	if res == 3600 || res == 14400 || res == 86400 || res == 300 || res == 60 || res == 900 {
		newRes = res
		return newRes
	}
	if res >= 86400 && res%86400 == 0 {
		return 86400
	}

	if res >= 14400 && res%14400 == 0 {
		return 14400
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
