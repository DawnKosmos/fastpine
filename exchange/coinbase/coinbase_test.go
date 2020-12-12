package coinbase

import (
	"fmt"
	"testing"
	"time"

	"github.com/dawnkosmos/fastpine/exchange"
)

func TestCoinbase(t *testing.T) {
	cb := New("", "")
	ch, err := cb.OHCLV("BTC-USD", 3600, exchange.DateToTime("01", "06", "2020"), time.Now().Unix())
	if err != nil {
		println(err.Error())
		return
	}

	f := ch[1].StartTime.Unix() - ch[0].StartTime.Unix()
	c1 := ch[0]
	for _, c := range ch[:] {
		fmt.Println(c.Close)
	}
}
