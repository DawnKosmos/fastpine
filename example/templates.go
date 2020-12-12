package main

import . "github.com/dawnkosmos/fastpine/series"

func CrossMA(src Series, ma func(src Series, l int) Series, signalLen int, invert bool) (buy Condition, sell Condition) {
	signal := ma(src, signalLen)
	if invert {
		sell = crossover(src, signal)
		buy = crossunder(src, signal)
	} else {
		buy = crossover(src, signal)
		sell = crossunder(src, signal)
	}
	return
}

func CrossDoubleMA(src Series, ma func(src Series, l int) Series, signalLen int, invert bool) (buy Condition, sell Condition) {
	signal := DoubleMA(ma, src, signalLen)
	if invert {
		sell = crossover(src, signal)
		buy = crossunder(src, signal)
	} else {
		buy = crossover(src, signal)
		sell = crossunder(src, signal)
	}
	return
}

func diamondTemplate(src, volume Series, ma func(src Series, l int) Series, l2 int, invert bool) (buy Condition, sell Condition) {
	outR := Sma(src, 2)
	outB1 := ma(outR, l2)
	outB2 := ma(outB1, l2)
	outB := SubF(outB1, outB2, 2.0)
	d := Sub(outR, outB)
	var diamond Series
	if volume == nil {
		diamond = Sma(d, 2)
	} else {
		diamond = Vwma(d, volume, 2)
	}
	if invert {
		sell, buy = momentumSwingBuySell(diamond)
	} else {
		buy, sell = momentumSwingBuySell(diamond)
	}

	return
}

func saphirTemplate(src, src2, volume Series, ma func(src Series, l int) Series, l2 int) (buy Condition, sell Condition) {
	outR1 := Sma(src, 2)
	outR2 := Sma(src, 2)
	outR := Add(outR1, outR2)
	outB1 := ma(outR, l2)
	outB2 := ma(outB1, l2)
	outB := SubF(outB1, outB2, 2.0)
	cc := Sub(outR, outB)

	var saph Series
	if volume == nil {
		saph = Sma(cc, 2)
	} else {
		cc1, volume1 := Offset(cc, 1), Offset(volume, 1)
		vwma1 := Add((MultF(cc, volume, 2)), Mult(cc1, volume1))
		saph = Div(vwma1, Add(Mult(volume, 2), volume1))
	}

	buy, sell = momentumSwingBuySell(saph)

	return

}
