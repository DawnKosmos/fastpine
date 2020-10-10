package series

import (
	"fmt"
	"testing"

	"github.com/dawnkosmos/fastpine/exchange"
	"github.com/dawnkosmos/fastpine/exchange/ftx"
)

func TestSeries(t *testing.T) {
	key := "aItoEjHCABXEuxg-C937wiOMZYMC2m0pr2pXYfQ3"
	secret := "v5wFE6jF599kA9y_agx9JegQqA_Xp_E45AFljlCZ"
	ug := exchange.NewUpdateGroup("Test", 60)
	f := ftx.New(key, secret, "", []string{""})
	src := NewOHCLV(f, "BTC-PERP", exchange.DateToTime("01", "06", "2019"), 3600*4, &ug)
	openclose := GetSRC(OC2, src)
	volume := GetSRC(VOLUME, src)
	//close := GetSRC(CLOSE, src)
	fmt.Println("parsing finished")
	diamond := Diamond(openclose, volume, 9, 4)
	c1, c2 := DiamondSignals(diamond)
	d := *diamond.Data()
	ss := *c2.DataB()
	for i, v := range *c1.DataB() {
		fmt.Println(d[i+2], v, ss[i])
	}
}

func Diamond(src Series, vol Series, lenRSI int, smooth int) Series {
	rsi := Sma(Rsi(src, lenRSI), 2)
	b1 := Sma(rsi, smooth)
	b2 := Sma(b1, smooth)
	bd := SubF(b1, b2, 2.0)
	d := Sub(rsi, bd)
	return Function(soos, 1, d, vol)
}

func soos(index int, v ...Series) float64 {
	src := v[0]
	vol := v[1]
	volume := vol.Value(index+0) + vol.Value(index+1)
	volX := vol.Value(index+0)*src.Value(index+0) + vol.Value(index+1)*src.Value(index+1)

	return volX / volume
}

func BuyCon(index int, v ...Series) (r float64) {
	src := v[0]
	v0, v1, v2 := src.Value(index+0), src.Value(index+1), src.Value(index+2)
	if v0 > v1 && v1 < v2 {
		r = 1
	} else {
		r = 0
	}
	return
}

func SellCon(index int, v ...Series) (r float64) {
	src := v[0]
	v0, v1, v2 := src.Value(index+0), src.Value(index+1), src.Value(index+2)
	if v0 < v1 && v1 > v2 {
		r = 1
	} else {
		r = 0
	}
	return
}

func DiamondSignals(d Series) (buy Condition, sell Condition) {
	buy = Function(BuyCon, 2, d)
	sell = Function(SellCon, 2, d)
	return
}
