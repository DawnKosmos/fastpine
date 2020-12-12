package series

import (
	"testing"

	"github.com/dawnkosmos/fastpine/exchange"
	"github.com/dawnkosmos/fastpine/exchange/coinbase"
)

func TestIff(t *testing.T) {
	ug := exchange.NewUpdateGroup("Group 1", 60)
	//b := bybit.New(false, nil, "", "")
	//d := deribit.New("", "")
	cb := coinbase.New("", "")
	chartdata := NewOHCLV(cb, "BTC-USD", exchange.DateToTime("01", "01", "2019"), 3600*48, &ug)
	close := Source(CLOSE, chartdata)
	open := Source(OPEN, chartdata)
	//volume := Source(VOLUME, chartdata)
	//high := Source(HIGH, chartdata)
	//low := Source(LOW, chartdata)

	dema := Ema(close, 24)
	PrintSeries(open, close, dema)
}
