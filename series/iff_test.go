package series

import (
	"testing"

	"github.com/dawnkosmos/fastpine/exchange"
	"github.com/dawnkosmos/fastpine/exchange/bybit"
)

func TestIff(t *testing.T) {
	ug := exchange.NewUpdateGroup("Group 1", 60)
	b := bybit.New(false, nil, "", "")
	chartdata := NewOHCLV(b, "BTCUSD", exchange.DateToTime("01", "01", "2020"), 3600*4, &ug)

	close := Source(CLOSE, chartdata)

	fastEma := Ema(close, 50)
	rsi := Rsi(close, 14)

	k := Iff(Greater(close, fastEma), fastEma, rsi)

	a := k.Data()

	Print
}
