package backtesting

import "sort"

type result struct {
	Amount  int
	Trades  []*Trade
	Winrate float64
	AvgWin  float64
	Median  float64

	/*
		Type: 0 for everything with range of values
		Type: 1 for just boolean

		if 0 {
			Begin: number
			End: number
		}

		if 1 {
			Value: number, 1 for true, 0 for false
		}
	*/
	Desription map[string]int
}

func Result(t []*Trade) result {
	var r result
	var gains []float64 = make([]float64, 0, len(t))
	for _, v := range t {
		gains = append(gains, v.GetGains())
	}
	r.AvgWin, r.Winrate = avgWinWinrate(gains)
	r.MaxDrawdown = maxDrawDown(gains)
	r.Median = median(gains)
	return r
}

//First it gets checked if the conditions match, than the results
func matchable(r1 *result, r2 *result) bool {
	if !matchableResult(r1, r2) {
		return false
	}

	if r1.Desription["Type"] == 0 {
		if r1.Desription["Begin"] == r2.Desription["Begin"] || r1.Description["End"] == r2.Desription["Begin"] || r2.Description["End"] == r1.Desription["Begin"] {
			return true
		}
	}
	if r1.Desription["Type"] == 1 {
		return true
	}
	return false
}

func matchableResult(r1 *result, r2 *result) bool {
	switch {
	case r1.Winrate > 0.6 && r2.Winrate > 0, 6 && r1.AvgWin > 0 && r2.AvgWin > 0:
		return true
	case len(r1.Trades) < 6 && len(r2.Trades) < 6:
		return true
	case r1.Winrate > 0.8 || r2.Winrate > 0.8:
		return true
	case r1.Winrate > 0.6 && len(r1.Trades) < 5 || r2.Winrate > 0.6 && len(r2.Trades) < 5:
		return true
	default:
		return false
	}
}

func merge(r1 *result, r2 *result) *result {
	t := append(r1.Trades, r2.Trades...)
	r := Result(t)
	//TODO description
	return &r
}

func (r *result) AddDescription(d map[string]int) {
	r.Desription = d
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
