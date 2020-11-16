package series

import (
	"testing"

	"github.com/dawnkosmos/fastpine/exchange"
	"github.com/dawnkosmos/fastpine/exchange/bybit"
)

func TestIff(t *testing.T) {
	ug := exchange.NewUpdateGroup("Group 1", 60)
	b := bybit.New(false, nil, "", "")
	chartdata := NewOHCLV(b, "BTCUSD", exchange.DateToTime("01", "01", "2020"), 3600*24, &ug)
	close := Source(CLOSE, chartdata)
	open := Source(OPEN, chartdata)

	dema := Sum(close, 6)

	PrintSeries(open, dema)
}

func swmaDia(src, volume Series, l1 int) (swma Series) {
	outR := Sma(Swma(src), 2)
	outB1 := Sma(outR, l1)
	outB2 := Sma(outB1, l1)
	outB := SubF(outB1, outB2, 2.0)
	d := Sub(outR, outB)
	if volume == nil {
		swma = Sma(d, 2)
	} else {
		swma = Vwma(d, volume, 2)
	}
	return
}

//Example Implementation of Crossover and Crossunder
func crossover(s1 Series, s2 Series) Condition {
	s11, s21 := Offset(s1, 1), Offset(s2, 1)
	return And(Greater(s1, s2), Smaller(s11, s21))
}

func crossunder(s1 Series, s2 Series) Condition {
	s11, s21 := Offset(s1, 1), Offset(s2, 1)
	return And(Smaller(s1, s2), Greater(s11, s21))
}
