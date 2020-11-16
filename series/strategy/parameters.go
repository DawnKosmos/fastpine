package strategy

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

/*Paremeters wants following input: ("parameterName", parameterValue)
So a STRING followed by a VALUE! following parameters are supported:
"fee", float
"balance": float
"pyramiding": int
"size": float
"slippage": float

Example:
Parameters("fee", 0.1, "balance" = 1000, "size", 0.5)
Adds a fee of 0.1% per trade, a balance of 1000, and a size of 0.5(50%) of the account size
*/
func (strat *Strategy) Parameters(p ...interface{}) error {
	if len(p)%2 != 0 {
		return errors.New("parameters input wrong, something is missing")
	}

	for i := 0; i < len(p); i = i + 2 {
		s, ok := p[i].(string)
		if !ok {
			return errors.New("Parameter Position:" + strconv.Itoa(i) + "is no string")
		}
		switch s {
		case "fee":
			err := strat.fee(p[i+1])
			if err != nil {
				return err
			}
		case "balance":
			err := strat.balance(p[i+1])
			if err != nil {
				return err
			}
		case "pyramiding":
			err := strat.pyramiding(p[i+1])
			if err != nil {
				return err
			}
		case "size":
			err := strat.size(p[i+1])
			if err != nil {
				return err
			}
		case "slippage":
			err := strat.slippage(p[i+1])
			if err != nil {
				return err
			}
		default:
			return errors.New("Not supported parameter " + s + " learn to read cunt")
		}
	}

	return nil
}

func (s *Strategy) fee(i interface{}) error {
	v, ok := i.(float64)
	if ok {
		s.Fee = v
		return nil
	}
	return errors.New("Wrong Type for Fee, I only accept float64")
}

func (s *Strategy) slippage(i interface{}) error {
	v, ok := i.(float64)
	if ok {
		s.Slippage = v
		return nil
	}
	d, ok := i.(int)
	if ok {
		s.Slippage = float64(d)
		return nil
	}

	return errors.New("Wrong Type for Slippage, I only accept float64")
}

func (s *Strategy) size(i interface{}) error {
	v, ok := i.(float64)
	if ok {
		s.Size = v
		return nil
	}
	return errors.New("Wrong Type for Size, I only accept float64")
}

func (s *Strategy) pyramiding(i interface{}) error {
	v, ok := i.(int)
	if ok {
		s.Pyramiding = v
		return nil
	}
	return errors.New("Wrong Type for Pyramiding, I only accept int")
}

func (s *Strategy) balance(i interface{}) error {
	v, ok := i.(float64)
	if ok {
		s.Balance = v
		return nil
	}
	d, ok := i.(int)
	if ok {
		s.Balance = float64(d)
		return nil
	}

	return errors.New("Wrong Type for Balance, I only accept float64")
}

func parameterShowTrades(i interface{}) bool {
	j, ok := i.(bool)
	if !ok {
		return false
	}
	return j
}

func parameterSide(t []Trades, i interface{}) []Trades {
	a, ok := i.(bool)
	if !ok || !a {
		return t
	}

	var out []Trades
	for _, v := range t {
		l, s := v.Longs(), v.Shorts()
		out = append(out, l, s)
	}
	return out
}

func parameterQty(qt interface{}) qtyType {
	i, ok := qt.(qtyType)
	if !ok {
		return PERCENT
	}
	return i
}

func parameterDescription(i interface{}) string {
	a, ok := i.(string)
	if !ok {
		return ""
	}
	return a
}

func parameterSplit(t []Trades, split interface{}) []Trades {
	//mc month Counter
	i, ok := split.(int)
	if !ok {
		fmt.Println("wrong input")
		return []Trades{}
	}

	var mc int
	//am actual month
	var out []Trades
	//tT temp Trades
	var tT Trades
	for _, vv := range t {
		var am time.Month = vv[0].EntryTime.Month()
		for _, v := range vv[1:] {
			vm := v.EntryTime.Month()
			if am != vm {
				am = vm
				mc++
			}
			if mc >= i {
				out = append(out, tT)
				mc = 0
				tT = Trades{}
			}

			tT = append(tT, v)
		}
	}
	return out
}
