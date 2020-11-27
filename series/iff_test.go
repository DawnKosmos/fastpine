package series

import (
	"testing"

	"github.com/dawnkosmos/fastpine/exchange"
	"github.com/dawnkosmos/fastpine/exchange/bybit"
)

func TestIff(t *testing.T) {
	ug := exchange.NewUpdateGroup("Group 1", 60)
	b := bybit.New(false, nil, "", "")
	chartdata := NewOHCLV(b, "BTCUSD", exchange.DateToTime("01", "01", "2019"), 3600*24, &ug)
	close := Source(CLOSE, chartdata)
	open := Source(OPEN, chartdata)
	//volume := Source(VOLUME, chartdata)
	//high := Source(HIGH, chartdata)
	//low := Source(LOW, chartdata)

	dema := DoubleMA(Ema, close, 30)
	PrintSeries(open, dema)
}
