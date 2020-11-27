package main

import (
	"fmt"

	"github.com/dawnkosmos/fastpine/exchange"
	"github.com/dawnkosmos/fastpine/exchange/ftx"
	. "github.com/dawnkosmos/fastpine/series"
)

func main() {
	/*	ug := exchange.NewUpdateGroup("Group 1", 60)
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



		//PrintSeries(open, close, signal, buy, sell)
	*/

	ug := exchange.NewUpdateGroup("Group 1", 60)
	f := ftx.New("", "", "", []string{""})

	//b := bybit.New(false, nil, "", "")
	fmt.Println("U WINNING?")
	chart := NewOHCLV(f, "ADA-PERP", exchange.DateToTime("01", "01", "2019"), 3600*3, &ug)
	//chart := NewOHCLV(b, "XRPEOS", exchange.DateToTime("01", "01", "2019"), 3600*6, &ug)

	mainDiamondIterations(chart)
	//time.Sleep(9 * time.Second)
	//mainDiamondIterations(chart)
}
