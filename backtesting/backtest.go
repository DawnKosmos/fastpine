package backtesting

import (
	"errors"

	"github.com/dawnkosmos/fastpine/series"
)

type Backtest struct {
	o      []*series.OHCLV
	env    []EnviromentFunc
	trades []AlgoFunc

	R Result
}

/*
func strategyTradeToTrade(ch Chart, trades []strategy.Trade) []Trade {
	cd := ch.ch
	var tr []Trade = make([]Trade, 0, len(trades))

	st := ch.Starttime - ch.Resolution
	var t Trade
	for _, v := range trades {
		t.EntryTime = int(v.EntryTime.Unix())
		t.ExitTime = int(v.ExitTime.Unix())
		t.EntryPrice = v.EntryPrice
		t.ExitPrice = v.ExitPrice
		t.Side = bool(v.Side)
		if t.EntryTime < cd[0].Timestamp {
			continue
		}
		t.EntryCondition = cd[(t.EntryTime-st)/ch.Resolution]
		t.ExitCondition = cd[(t.ExitTime-st)/ch.Resolution]
		tr = append(tr, t)
	}

	return tr
}*/

type Algo struct {
	description string
	longs       []Trade
	shorts      []Trade
}

type AlgoFunc func(o *series.OHCLV) (des string, longs []series.Trade, shorts []series.Trade)
type EnviromentFunc func(o *series.OHCLV) (des string, evn series.Condition)

func New(o *series.OHCLV, trades []strategy.Trade, IndicatorLayout []string, indicators ...series.Series) *Backtest {
	var b Backtest
	b.e = o.E

}

func (b *Backtest) Candlestickdata(o ...*series.OHCLV) error {
	_ = o
	return errors.New("kek")
}

func (b *Backtest) Trades(a ...AlgoFunc) error {
	_ = t
	return errors.New("kek")
}

func (b *Backtest) Indicators(IndicatorLayout []string, s ...series.Series) {
	_ = "kek"
}

/*
token := {"Compare","ID", "ID | Number"} | {"Value", "ID", "Int"},
Expression := token | {token, "or | and", token}}
*/
func (b *Backtest) Condition(con [][]string) {

}

func (b *Backtest) Enviroment(env EnviromentFunc) {

}
