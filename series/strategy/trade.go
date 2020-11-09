package strategy

import (
	"errors"
	"fmt"
	"time"

	"github.com/dawnkosmos/fastpine/exchange"
)

type Side bool

const (
	LONG  Side = true
	SHORT Side = false
)

type qtyType int

const (
	PERCENT qtyType = 0
	USD     qtyType = 1
)

type Trade struct {
	StoppedOut bool
	Side       Side // true = Long, false = Short; gets translated when outputing

	EntryPrice float64
	ExitPrice  float64

	EntryTime time.Time
	ExitTime  time.Time

	Size float64
}

func (t *Trade) GetGains(fee float64, slippage float64) float64 {
	var x float64
	if t.Side {
		x = (t.ExitPrice-slippage-t.EntryPrice+slippage)/t.EntryPrice + slippage
	} else {
		x = -1*(t.ExitPrice-slippage-t.EntryPrice+slippage)/t.EntryPrice + slippage
	}
	return x*t.Size - (fee * 0.01)
}

func CreateTrade(side Side, Size float64, entry, exit exchange.Candle) (Trade, error) {
	if entry.StartTime.Unix() == exit.StartTime.Unix() {
		return Trade{}, errors.New("Same Candle")
	}
	var t Trade = Trade{
		Side:       side,
		EntryPrice: entry.Open,
		ExitPrice:  exit.Open,
		EntryTime:  entry.StartTime,
		ExitTime:   exit.StartTime,
		Size:       Size,
	}
	return t, nil
}

func CreateTradeComplicated() {
	_ = "kek"
}

func (t Trade) String() string {
	var side string
	if t.Side {
		side = "Buy"
	} else {
		side = "Sell"
	}
	return fmt.Sprintf("%s Entrytime: %v - Exittime: %v - Entryprice: %f - Exitprice: %f", side, t.EntryTime.Format(time.RFC822), t.ExitTime.Format(time.RFC822), t.EntryPrice, t.ExitPrice)
}

//TRADES
type Trades []Trade

func (t Trades) Less(i, j int) bool {
	return t[i].EntryTime.Unix() < t[j].EntryTime.Unix()
}

func (t Trades) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Trades) Len() int {
	return len(t)
}

func (t Trades) Longs() Trades {
	new := make([]Trade, 0, len(t)/2)
	for _, v := range t {
		if v.Side {
			new = append(new, v)
		}
	}
	return new
}

func (t Trades) Shorts() Trades {
	new := make([]Trade, 0, len(t)/2)
	for _, v := range t {
		if !v.Side {
			new = append(new, v)
		}
	}
	return new
}

func (t Trades) TradeGains(fee float64, slippage float64) []float64 {
	gains := make([]float64, 0, t.Len())
	for _, v := range t {
		gains = append(gains, v.GetGains(fee, slippage))
	}

	return gains
}

func (s *Strategy) GetTrades(b ...bool) Trades {
	if len(b) == 0 {
		return s.trades
	}

	if b[0] {
		return s.trades.Longs()
	}

	return s.trades.Shorts()

}
