package series

import (
	"testing"

	"github.com/dawnkosmos/fastpine/exchange"
	"github.com/dawnkosmos/fastpine/exchange/deribit"
)

func TestIff(t *testing.T) {
	ug := exchange.NewUpdateGroup("Group 1", 60)
	//b := bybit.New(false, nil, "", "")
	//d := deribit.New("", "")
	cb := deribit.New("", "")
	chart := NewOHCLV(cb, "BTC-PERPETUAL", exchange.DateToTime("01", "01", "2019"), 3600*4, &ug)

	volume := Source(VOLUME, chart)
	//high := Source(HIGH, chartdata)
	low := Source(LOW, chart)

	src := Sma(HCL3(chart), 2)
	outR := Sma(src, 6)
	outB1 := Sma(outR, 6)
	outB2 := Sma(outB1, 6)
	outB := SubF(outB1, outB2, 2.0)
	d := Sub(outR, outB)
	diamond := Vwma(d, volume, 2)
	buy, _ := momentumSwingBuySell(diamond)

	dema := LongDiv(low, diamond, buy)
	PrintSeries(low, diamond, buy, dema)
}

func momentumSwingBuySell(d Series) (buy Condition, sell Condition) {
	d1, d2 := Offset(d, 1), Offset(d, 2)
	buy = And(Greater(d, d1), Smaller(d1, d2))
	sell = And(Smaller(d, d1), Greater(d1, d2))
	return
}
