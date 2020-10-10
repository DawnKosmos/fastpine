package series

/*
func TestRSI(t *testing.T) {
	key := "aItoEjHCABXEuxg-C937wiOMZYMC2m0pr2pXYfQ3"
	secret := "v5wFE6jF599kA9y_agx9JegQqA_Xp_E45AFljlCZ"
	ug := exchange.NewUpdateGroup("Test", 60)
	f := ftx.New(key, secret, "", []string{""})
	src := NewOHCLV(f, "BTC-PERP", exchange.DateToTime("01", "01", "2020"), 86400, &ug)
	close := GetSRC(CLOSE, src)
	/*rsi := Rsi(close, 14)
	c := close.Data()
	r := rsi.Data()
	sma := Sma(rsi, 20)
	addRS := Add(sma, rsi)
	addRSS := Add(addRS)
	a := addRS.Data()*/

//openclose := GetSRC(OC2, src)

/*
	rsi = sma(rsi(open + close, 10), 2)
	b1 = sma(rsi, 4)
	b2 = sma(b2, 4)
	bd = 2 * b1 - b2

	d = rsi - bd
	diamond = sma(d, 2)
*/

/*rsi := Sma(Rsi(openclose, 10), 2)
	b1 := Sma(rsi, 4)
	b2 := Sma(b1, 4)
	bd := SubF(b1, b2, 2.0)
	d := Sub(rsi, bd)
	diamond := Sma(d, 2)



	macd, signal, hist := macd(close, 12, 26, 9)

	m := macd.Data()
	s := signal.Data()
	h := hist.Data()
	for i, v := range *s {
		fmt.Println((*m)[i+8], v, (*h)[i])
	}
}

func macd(src Series, fast int, slow int, signal int) (macd Series, sig Series, hist Series) {
	f := Ema(src, fast)
	s := Ema(src, slow)
	macd = Sub(f, s)
	sig = Ema(macd, signal)
	hist = Sub(macd, sig)
	return
}

/*
btdRSI = sma(rsi(open + close, 10), 2)
btdB1 = sma(btdRSI, 4)
btdB2 = sma(btdB1, 4)
btdB = 2 * btdB1 - btdB2

diffBtd = btdRSI - btdB
sma_2 = sma(diffBtd, 2)
*/
