package deribit

import (
	"fmt"
	"testing"
	"time"

	"github.com/dawnkosmos/fastpine/exchange"
)

func TestJa(t *testing.T) {
	d := New("", "")
	ch, err := d.OHCLV("BTC-PERPETUAL", 3600*4, exchange.DateToTime("01", "01", "2019"), time.Now().Unix())
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, c := range ch[1:] {
		if c.StartTime.Unix()-ch[i].StartTime.Unix() != 3600*4 {
			fmt.Println(i, "U suck coding")
		}
	}
}
