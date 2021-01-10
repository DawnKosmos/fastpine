package main

import (
	"fmt"

	"github.com/dawnkosmos/fastpine/exchange"
	"github.com/dawnkosmos/fastpine/exchange/deribit"

	. "github.com/dawnkosmos/fastpine/series"
)

func main() {
	/*ug := exchange.NewUpdateGroup("Group 1", 60)
	//b := bybit.New(false, nil, "", "")
	f := ftx.New("", "", "", []string{""})

	//Strategies := []Series{Rsi(openclose, 6+i), Ema(openclose, 6+i), Roc(openclose, 4+i), dai(open, high, low, close, volume, 7, 2+i), Swma(openclose)}

	chartdata := NewOHCLV(f, "ETH-PERP", exchange.DateToTime("01", "01", "2019"), 3600*4, &ug)

	//close := Source(CLOSE, chartdata)
	openclose := Source(OC2, chartdata)
	//volume := Source(VOLUME, chartdata)
	buy, sell := diamondTemplate(Swma(openclose), nil, Sma, 4+2, false)

	aa := strategy.NewSimple(chartdata, buy, sell)
	_, in := aa.Results("side", false, "trades", false)

	fmt.Println(in[0]) */
	//PrintSeries(open, close, signal, buy, sell)

	ug := exchange.NewUpdateGroup("Group 1", 60)
	//f := ftx.New("", "", "", []string{""})

	//b := bybit.New(false, nil, "", "")
	cb := deribit.New("", "")
	fmt.Println("U WINNING Son?")
	chart := NewOHCLV(cb, "ETH-PERPETUAL", exchange.DateToTime("01", "11", "2018"), 3600*3, &ug)
	mainDiamondIterations(chart)

}
