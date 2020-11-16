package main

import . "github.com/dawnkosmos/fastpine/series"

func diamond(src Series, volume Series, l1, l2 int) (diamond Series) {
	outR := Sma(Rsi(src, l1), 2)
	outB1 := Sma(outR, l2)
	outB2 := Sma(outB1, l2)
	outB := SubF(outB1, outB2, 2.0)
	d := Sub(outR, outB)
	if volume == nil {
		diamond = Sma(d, 2)
	} else {
		diamond = Vwma(d, volume, 2)
	}
	return
}

func saphir(src1, src2, volume Series, l1, l2 int) (saphir Series) {
	r1 := Sma(Rsi(src1, l1), 2)
	r2 := Sma(Rsi(src2, l2), 2)
	r := AddF(r1, r2, 2)
	b1 := Sma(r, l2)
	b2 := Sma(b1, l2)
	b := SubF(b1, b2, 2)
	cc := Sub(r, b)
	if volume == nil {
		saphir = Sma(cc, 2)
	} else {
		cc1, volume1 := Offset(cc, 1), Offset(volume, 1)
		vwma1 := Add((MultF(cc, volume, 2)), Mult(cc1, volume1))
		saphir = Div(vwma1, Add(Mult(volume, 2), volume1))
	}
	return
}

func swmaDia(src, volume Series, l1 int) (swma Series) {
	outR := Sma(Swma(src), 2)
	outB1 := Sma(outR, l1)
	outB2 := Sma(outB1, l1)
	outB := SubF(outB1, outB2, 2.0)
	d := Sub(outR, outB)
	if volume == nil {
		swma = Sma(d, 2)
	} else {
		swma = Vwma(d, volume, 2)
	}
	return
}

func dai(open, high, low, close, volume Series, l1, l2 int) Series {
	sh := Vwma(Sum(Add(high, close), l1), volume, l2)
	lo := Vwma(Sum(Sub(close, low), l1), volume, l2)
	return Div(Sub(lo, sh), sh)
}

func momentumSwingBuySell(d Series) (buy Condition, sell Condition) {
	d1, d2 := Offset(d, 1), Offset(d, 2)
	buy = And(Greater(d, d1), Smaller(d1, d2))
	sell = And(Smaller(d, d1), Greater(d1, d2))
	return
}

func longCon(prozent float64, len int, close Series, low Series) Condition {
	lowest := Lowest(low, len)
	r1 := Sub(Div(close, lowest), 1)
	return Greater(r1, prozent/100)
}

func shortCon(prozent float64, len int, close, high Series) Condition {
	highest := Highest(high, len)
	r1 := Sub(Div(close, highest), 1)
	return Smaller(r1, -prozent/100)
}

//Example Implementation of Crossover and Crossunder
func crossover(s1 Series, s2 Series) Condition {
	s11, s21 := Offset(s1, 1), Offset(s2, 1)
	return And(GreaterEqual(s1, s2), Smaller(s11, s21))
}

func crossunder(s1 Series, s2 Series) Condition {
	s11, s21 := Offset(s1, 1), Offset(s2, 1)
	return And(SmallerEqual(s1, s2), Greater(s11, s21))
}

func doubleEMA(s Series, l int) Series {
	e := Ema(s, l)
	return SubF(e, Ema(e, l), 2)
}

func doubleSMA(s Series, l int) Series {
	e := Sma(s, l)
	return SubF(e, Sma(e, l), 2)
}

func doubleRMA(s Series, l int) Series {
	e := Rma(s, l)
	return SubF(e, Rma(e, l), 2)
}
