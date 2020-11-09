package strategy

import (
	"errors"
	"fmt"
	"strconv"
)

/*
Results
interface parameters:
"split" - 1-12 get the results in monthly pakets |||
"qty" - PERCENT | USD Trade a % of your equity or always same amount of USD |||
"trades" - true | false  Print all the trades |||
"side" - true | false  Seperates result by Long/Short
*/
func (s *Strategy) Results(parameters ...interface{}) error {
	p := parameters
	if len(p)%2 != 0 {
		return errors.New("parameters input wrong, something is missing")
	}

	//default Parameters
	qt := PERCENT
	trades := []Trades{s.GetTrades()}
	printTrades := false

	var sideSplit bool

	for i := 0; i < len(p); i = i + 2 {
		str, ok := p[i].(string)
		if !ok {
			return errors.New("Parameter Position:" + strconv.Itoa(i) + "is no string")
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
		}
	}

	result(trades, qt, sideSplit, printTrades, s.Balance, s.Fee, s.Slippage)
	return nil
}

func result(t []Trades, qt qtyType, sideSplit bool, printTrades bool, balance, fee, slippage float64) {
	for _, v := range t {
		instance(v, qt, sideSplit, printTrades, balance, fee, slippage)
	}
}

func instance(t Trades, qt qtyType, ss bool, pt bool, balance, fee, slippage float64) {
	startingMonth, endingMonth := t[0].EntryTime.Month(), t[len(t)-1].EntryTime.Month()
	side := "Long & Short"
	amount := len(t)

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

	str := fmt.Sprintf("Side: %s, Trades: %d, Pnl: %.1f, Winrate: %.2f, Start: %v, End: %v", side, amount, pnl, winrate, startingMonth, endingMonth)
	fmt.Println(str)

	if pt {
		fmt.Println("-----------------------------------------------------")
		for _, v := range t {
			fmt.Println(v, v.GetGains(fee, slippage)*100, "%")
		}
		fmt.Println("-----------------------------------------------------")
	}
}
