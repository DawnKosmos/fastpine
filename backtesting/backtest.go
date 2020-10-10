package backtesting

import "sort"

type Backtest struct {
	Exchage string
	Ticker  string
	Trades  []Trade
	R       Result
}

type Result struct {
	Amount      int
	Winrate     float64
	AvgWin      float64
	Median      float64
	MaxDrawdown float64
}

func New(exchange string, ticker string, trades *[]Trade) *Backtest {
	var b Backtest
	b.Exchage, b.Ticker = exchange, ticker
	b.Trades = *trades
	b.R = getResult(b.Trades)
	return &b
}

func getResult(t []Trade) Result {
	var r Result
	var gains []float64 = make([]float64, 0, len(t))
	for _, v := range t {
		gains = append(gains, v.GetGains())
	}
	r.AvgWin, r.Winrate = avgWinWinrate(gains)
	r.MaxDrawdown = maxDrawDown(gains)
	r.Median = median(gains)

	return r

}

func avgWinWinrate(f []float64) (float64, float64) {
	var gains float64
	var count int = 0
	for _, a := range f {
		gains += a
		if a > 0 {
			count++
		}
	}
	return gains / float64(len(f)), float64(count) / float64(len(f)) * 100
}

func maxDrawDown(f []float64) float64 {
	var maxDD float64 = 1.0
	var actualDD float64 = 1.0
	var maxGains float64 = 1.0
	for _, v := range f {
		if v < 0 {
			actualDD = actualDD * (1 + v)
			if actualDD < maxDD {
				maxDD = actualDD
			}
		} else {
			actualDD = 1
		}
		maxGains *= (1 + v)
	}
	return maxGains
}

func median(f []float64) (m float64) {
	l := len(f)
	sort.Sort(sort.Float64Slice(f))
	k := l / 2
	if l%2 == 0 {
		m = (f[k-1] + f[k]) / 2
	} else {
		m = f[k]
	}
	return
}
