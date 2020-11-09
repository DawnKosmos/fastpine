package backtesting

import (
	"strconv"

	"github.com/dawnkosmos/fastpine/series"
)

type Simp struct {
	ch         Chart
	trades     [][]*Trade
	conditions []Cons
	Layout     map[string]int
}

func Simple(o *series.OHCLV, IndicatorLayout []string, Indicators ...series.Series) *Simp {
	var s Simp

	//init chart
	var ch Chart
	ch.Exchange = (*o.E).Name()
	ch.Resolution = o.Resolution()
	ch.Ticker = o.Ticker
	ch.Starttime = int(o.Starttime())
	ch.src = exchangeCandleToCandle(o.Data(), Indicators...)

	//Init Simple
	s.ch = ch
	layout := make(map[string]int)
	for i, v := range IndicatorLayout {
		layout[v] = i
	}
	s.Layout = layout
	return &s
}

func (s *Simp) Strategy(t ...[]series.Trade) {
	for _, v := range t {
		s.trades = append(s.trades, seriesTradeToTrade(s.ch, v))
	}
}

//Comp, CompV
//Part
func (s *Simp) SetCondition(sc [][]string) {
	for _, v := range sc {
		switch v[0] {
		case "gr", "Gr", "greater", "Greater":
			b := stringtToComp(s.Layout, v[1], v[2])
		case "Part", "part":
		}
	}
}

func stringtToComp(layout map[string]int, s1, s2 string) Backtest {
	n1, err := strconv.ParseFloat(s1, 64)
	if err == nil {
		_, err := strconv.ParseFloat(s2)
		if err != nil {
			return nil
		}
		return CompV(false, layout[s2], n1, s2, s1)
	}
	p1, ok := layout[s1]
	if !ok {
		return nil
	}

	n2, err := strconv.ParseFloat(s2, 64)
	if err == nil {
		_, err := strconv.ParseFloat(s)
		if err != nil {
			return nil
		}
		return CompV(true, p1, n2, s1, s2)
	}

	p2, ok := layout[s2]
	if !ok {
		return nil
	}

	return Comp(p1, p2, s1, s2)
}


func stringToPart{layout map[string]int, s1,s2 string} Backtest{
	
}