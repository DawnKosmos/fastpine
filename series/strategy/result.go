package strategy

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"
)

/*
Results
interface parameters:
"split" - 1-12 get the results in monthly pakets |||
"qty" - PERCENT | USD Trade a % of your equity or always same amount of USD |||
"trades" - true | false  Print all the trades |||
"side" - true | false  Seperates result by Long/Short
*/
func (s *Strategy) Results(parameters ...interface{}) (error, []Instance) {
	p := parameters
	if len(p)%2 != 0 {
		return errors.New("parameters input wrong, something is missing"), nil
	}
	var description string

	//default Parameters
	qt := PERCENT
	trades := []Trades{s.GetTrades()}
	printTrades := false

	var sideSplit bool

	for i := 0; i < len(p); i = i + 2 {
		str, ok := p[i].(string)
		if !ok {
			return errors.New("Parameter Position:" + strconv.Itoa(i) + "is no string"), nil
		}
		switch str {
		case "split":
			trades = parameterSplit(trades, p[i+1])
		case "qty":
			qt = parameterQty(p[i+1])
		case "trades":
			printTrades = parameterShowTrades(p[i+1])
		case "side":
			trades = parameterSide(trades, p[i+1])
			sideSplit = true
		case "description":
			description = parameterDescription(p[i+1])
		}
	}

	return nil, result(trades, qt, sideSplit, printTrades, s.Balance, s.Fee, s.Slippage, description)
}

func result(t []Trades, qt qtyType, sideSplit bool, printTrades bool, balance, fee, slippage float64, description string) []Instance {
	var ii []Instance

	for _, v := range t {
		ii = append(ii, instance(v, qt, sideSplit, printTrades, balance, fee, slippage, description))
	}
	return ii
}

type Instance struct {
	description string
	startMonth  time.Time
	endMonth    time.Time
	side        string

	fee      float64
	slippage float64

	trades  Trades
	winrate float64
	pnl     float64
	avgWin  float64

	pt bool
}

func (s Instance) Description() string {
	return s.description
}

func (s Instance) Print() {

	str := fmt.Sprintf("Side: %s, Trades: %d, Pnl: %.1f, Winrate: %.2f, AvgGain %f.3f, Start: %v %v End: %v %v", s.side, len(s.trades), s.pnl, s.winrate, s.avgWin, s.startMonth.Month(), s.startMonth.Year(), s.endMonth.Month(), s.endMonth.Year())
	fmt.Println(str)

	if s.pt {
		fmt.Println("-----------------------------------------------------")
		for _, v := range s.trades {
			fmt.Println(v, v.GetGains(s.fee, s.slippage)*100, "%")
		}
		fmt.Println("-----------------------------------------------------")
	}

}

func instance(t Trades, qt qtyType, ss bool, pt bool, balance, fee, slippage float64, description string) Instance {
	if len(t) == 0 {
		return Instance{
			description: "",
			endMonth:    time.Now(),
			startMonth:  time.Now(),
			winrate:     0.0,
			pnl:         0.0,
			slippage:    0.0,
			fee:         0.0,
			trades:      t,
			pt:          pt,
			side:        "0",
		}
	}

	startingMonth, endingMonth := t[0].EntryTime, t[len(t)-1].EntryTime
	side := "Long & Short"

	if ss {
		if t[0].Side {
			side = "Long "
		} else {
			side = "Short"
		}
	}

	gainz := t.TradeGains(fee, slippage)
	winrate := Winrate(gainz)

	var pnl float64 = balance

	if qt == 0 {
		pnl *= Product(gainz)
	} else {
		for _, v := range gainz {
			pnl += balance * (1.0 + v)
		}
	}

	avgWin := math.Pow(pnl, 1.0/float64(len(t)))

	inst := Instance{
		description: description,
		endMonth:    endingMonth,
		startMonth:  startingMonth,
		winrate:     winrate,
		pnl:         pnl,
		slippage:    slippage,
		fee:         fee,
		trades:      t,
		pt:          pt,
		side:        side,
		avgWin:      avgWin,
	}
	return inst
}

type Instances []Instance

func (a Instances) Len() int           { return len(a) }
func (a Instances) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Instances) Less(i, j int) bool { return a[i].avgWin < a[j].avgWin }
