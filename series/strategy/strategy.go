package strategy

import (
	"fmt"
	"sort"

	"github.com/dawnkosmos/fastpine/exchange"
	"github.com/dawnkosmos/fastpine/helper"
	"github.com/dawnkosmos/fastpine/series"
)

type Strategy struct {
	ch *series.OHCLV

	Balance    float64
	Size       float64
	Fee        float64
	Slippage   float64
	Pyramiding int

	stopLong, stopShort series.Series

	trades Trades
}

/*NewSimple creates a simple strategy but you can't add stops or have in candle execution(breakout Trade)
Parameters are set as Pairs at the end of the functino: 1.) "parameterName" as string, followed by the parameterValue as float64|int
e.g NewSimple(ohcvl, buy, sell, "fee", 0.01, "slippage": 0.5, "balance": 300, "size", 0.6)
Parameters are following: balance, fee, size, pyramiding, slippage*/
func NewSimple(o *series.OHCLV, buy series.Condition, sell series.Condition, parameters ...interface{}) *Strategy {
	//short Slices so all got same lenght to simplify iterating
	//chartdata shifted 1 more, due trades get executed the next candle after a signal
	ch, l, s := o.Data(), buy.DataB(), sell.DataB()
	sl := helper.Min(len(ch), len(l), len(s))

	ch = ch[len(ch)-sl+1:]
	l = l[len(l)-sl:]
	s = s[len(s)-sl:]

	//Init
	strat := Strategy{
		ch:         o,
		Balance:    1,
		Fee:        0,
		Pyramiding: 1,
		Size:       1.0,
	}

	err := strat.Parameters(parameters...)
	if err != nil {
		fmt.Println(err.Error())
	}

	var trades []Trade
	var t Trade
	var tempOrderLong, tempOrderShort []exchange.Candle
	p := strat.Pyramiding

	for i, c := range ch {
		if l[i] {
			for i := 0; i < helper.Min(len(tempOrderShort), p); i++ {
				t, err = CreateTrade(SHORT, strat.Size, tempOrderShort[i], c)
				if err != nil {
					fmt.Println("Get Shorts at", i, err)
					continue
				}
				trades = append(trades, t)
			}
			tempOrderShort = tempOrderShort[:0]
			tempOrderLong = append(tempOrderLong, c)
		}
		if s[i] {
			for i := 0; i < helper.Min(len(tempOrderLong), p); i++ {
				t, err = CreateTrade(LONG, strat.Size, tempOrderLong[i], c)
				if err != nil {
					fmt.Println("Get Longs at", i, err)
					continue
				}
				trades = append(trades, t)
			}
			tempOrderLong = tempOrderLong[:0]
			tempOrderShort = append(tempOrderShort, c)
		}
	}

	sort.Sort(Trades(trades))
	strat.trades = trades
	return &strat
}
