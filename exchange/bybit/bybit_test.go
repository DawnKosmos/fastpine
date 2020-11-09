package bybit

import (
	"fmt"
	"testing"
	"time"

	"github.com/dawnkosmos/fastpine/exchange"
)

func TestChartdata(t *testing.T) {
	b := New(false, nil, "", "")
	ch, err := b.OHCLV("BTCUSD", 3600*6, exchange.DateToTime("01", "02", "2020"), time.Now().Unix())
	if err != nil {
		println(err.Error())
	}

	f := ch[1].StartTime.Unix() - ch[0].StartTime.Unix()
	c1 := ch[0]
	for _, c := range ch[2:] {
		if c.StartTime.Unix()-c1.StartTime.Unix() == f {
			fmt.Print(1)
		} else {
			fmt.Print(c.StartTime.Unix(), c1.StartTime.Unix())
		}
		c1 = c
	}
}
