package main

import (
	"github.com/dawnkosmos/fastpine/exchange"
	"github.com/dawnkosmos/fastpine/exchange/bybit"
	. "github.com/dawnkosmos/fastpine/series"
	"github.com/dawnkosmos/fastpine/series/strategy"
)

func main() {
	//An updategroup is later needed for live trading, but not implemented right now
	ug := exchange.NewUpdateGroup("Group 1", 60)
	//Source init
	//Right now I barely support Bybit and FTX
	b := bybit.New(false, nil, "", "")
	/*Chartdata always downloads the data from the exchange
	I will add support that it saves the data either in a .txt, .csv or a sql database
	Resolution is in seconds and supports most
	*/
	chartdata := NewOHCLV(b, "BTCUSD", exchange.DateToTime("01", "01", "2020"), 3600*4, &ug)

	//Getting the needed Source
	close := Source(CLOSE, chartdata)

	//slowEma = ema(close, 200)
	//fastEma = ema(close, 50)
	//goldenCross = crossover(fastEma,slowEma)
	//deathCross = crossunder(fastEma,slowEma)
	slowEma := Ema(close, 200)
	fastEma := Ema(close, 50)
	goldenCross := crossover(fastEma, slowEma)
	deathCross := crossunder(fastEma, slowEma)

	PrintSeries(close, fastEma, slowEma, goldenCross, deathCross)

	//Running the strategy, with a fee of 0.07% size of 100% and a starting balance of 100
	s := strategy.NewSimple(chartdata, goldenCross, deathCross, "fee", 0.07, "size", 1.0, "balance", 100)
	_ = s.Results("side", true, "trades", true)

}

//Example Implementation of Crossover and Crossunder
func crossover(s1 Series, s2 Series) Condition {
	// s11 = s1[1]
	// s21 = s2[1]
	s11, s21 := Offset(s1, 1), Offset(s2, 1)

	// s1>s2, s1 < s2
	return And(Greater(s1, s2), Smaller(s11, s21))
}

func crossunder(s1 Series, s2 Series) Condition {
	// s11 = s1[1]
	// s21 = s2[1]
	s11, s21 := Offset(s1, 1), Offset(s2, 1)
	// s1 < s2, s1 > s2
	return And(Smaller(s1, s2), Greater(s11, s21))
}
