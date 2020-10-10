package main

import (
	"github.com/dawnkosmos/fastpine/exchange"
)


type Backtest int

const(
	TRADING Backtest = iota
	DATA
	ALERTS
	BACKTEST
)

func main(){
	e := exchange.New("FTX")
	i := e.NewInstance(e, "BTC-PERP", 3600, Frequenz)
	heikinAshi := heikinAshi(i)
	ema := s.ema(heikinAshi.Close(), 5)
	sma := s.ema(heikinAshi.Open(),6)
	k := offset(s.add(ema,sma),10)

	source := exchange.New("FTX","BTC-PERP", Resolution, starttime, cycletime, TRADING)
	source2 := heikinAshi(source)
	source.ADDAccount("key","secret", "")
	ema := EMA(source.Close, 5)
	sma := EMA(source.Open, 5)

	k := OFFSET(ADD(ema, sma), 1)

	r := RSI(k, 14)

	up := GREATHERTHAN(r, 20)
	NicerDicer := IFF(up, k, r)
	
	source.Buy("MARKET","Candle Close", up, "40%")
	source.Stop("Sell"),"Candle Close", up, "20%"
	source.Sell("Market","Candle Close", down)

}




func MACD(src *Series, slow,fast,signal int)(*Series,*Series,*Series){
	slowE := EMA(src, slow)
	fastE := EMA(src, fast)

	macd := SUB(slow,fastE)

	signal := ema(macd, signal)

	histo := SUB(macd, signal)

	return macd, signal, histo
}