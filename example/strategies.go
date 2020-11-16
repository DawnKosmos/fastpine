package main

import (
	"fmt"
	"sort"
	"time"

	. "github.com/dawnkosmos/fastpine/series"
	"github.com/dawnkosmos/fastpine/series/strategy"
)

func mainMaCrossIterating(chart *OHCLV) {
	open := Source(OPEN, chart)
	close := Source(CLOSE, chart)
	low := Source(LOW, chart)
	high := Source(HIGH, chart)
	volume := Source(VOLUME, chart)
	openclose := Source(OC2, chart)
	//ohclv := Source(OHCL4, chart)
	//hl2 := Source(HL2, chart)

	swma := func(s Series, l int) Series {
		_ = l
		return Swma(s)
	}

	maFuncs := []func(src Series, l1 int) Series{
		Ema,
		Sma,
		Rma,
		swma,
	}
	trades := 0
	nau := time.Now()

	var srcAlgo string
	var maAlgo string

	var inst strategy.Instances

	//DaiCross RSICross
	fmt.Println("Start Iterating")

	for i := 0; i < 6; i++ {
		Strategies := []Series{dai(open, high, low, close, volume, 6+i, 2), swmaDia(openclose, nil, 3+i), Swma(openclose), Rsi(openclose, 10+i)}
		for ss, s := range Strategies {
			switch ss {
			case 0:
				srcAlgo = "dai"
			case 1:
				srcAlgo = "swmaDia"
			case 2:
				srcAlgo = "swma"
			case 3:
				srcAlgo = "RSI"
			}

			for ee, e := range maFuncs {
				for k := 0; k < 7; k++ {
					buy, sell := CrossDoubleMA(s, e, k+5, false)
					strat := strategy.NewSimple(chart, buy, sell)
					switch ee {
					case 0:
						maAlgo = "EMA"
					case 1:
						maAlgo = "SMA"
					case 2:
						maAlgo = "RMA"
					case 3:
						maAlgo = "SWMA"
					}
					de := fmt.Sprintf("%d | %s | %s | %d", i, srcAlgo, maAlgo, k+5)
					_, ii := strat.Results("side", false, "description", de)
					inst = append(inst, ii...)
					trades++

				}
			}

		}
	}
	sort.Sort(inst)
	fmt.Println(trades, "Backtest calculated in ", time.Since(nau))

	for _, i := range inst {
		fmt.Println(i.Description())
		i.Print()
	}
}

func mainDiamondIterations(chart *OHCLV) {
	open := Source(OPEN, chart)
	close := Source(CLOSE, chart)
	low := Source(LOW, chart)
	high := Source(HIGH, chart)
	volume := Source(VOLUME, chart)
	openclose := Source(OC2, chart)
	//ohclv := Source(OHCL4, chart)
	//hl2 := Source(HL2, chart)

	vwma := func(s Series, l int) Series {
		return Vwma(s, volume, l)
	}

	swma := func(s Series, l int) Series {
		_ = l
		return Swma(s)
	}

	maFuncs := []func(src Series, l1 int) Series{
		Ema,
		Sma,
		Rma,
		doubleEMA,
		doubleSMA,
		doubleRMA,
		vwma,
		swma,
	}
	trades := 0
	backtest := 0
	nau := time.Now()

	var srcAlgo string
	var maAlgo string

	var inst strategy.Instances

	//DaiCross RSICross
	fmt.Println("Start Iterating")
	/*fish := Fisher(hl2, 9)
	Var conAlgo string
	cons := []Condition{Greater(Ema(close, 21), close), Greater(fish, Offset(fish, 1)), Greater(Ema(close, 50), Ema(close, 200)), Greater(fish, 0), Greater(Rsi(close, 14), 60)}

	for cc, c := range cons[:4] {
		switch cc {
		case 0:
			conAlgo = "21EMA"
		case 1:
			conAlgo = "Fish+"
		case 2:
			conAlgo = "GoldenBull"
		case 3:
			conAlgo = "Fish>0"
		case 4:
			conAlgo = "Rsi>60"
		}
	*/
	for i := 0; i < 7; i++ {
		Strategies := []Series{Rsi(openclose, 6+i), Ema(openclose, 6+i), Roc(openclose, 4+i), dai(open, high, low, close, volume, 7, 2+i), Swma(openclose)}
		for ss, s := range Strategies {
			switch ss {
			case 0:
				srcAlgo = "RSI"
			case 1:
				srcAlgo = "EMA"
			case 2:
				srcAlgo = "Roc"
			case 3:
				srcAlgo = "DAI"
			case 4:
				srcAlgo = "Swma"
			}
			for ee, e := range maFuncs {
				switch ee {
				case 0:
					maAlgo = "EMA"
				case 1:
					maAlgo = "SMA"
				case 2:
					maAlgo = "RMA"
				case 3:
					maAlgo = "DoubleEma"
				case 4:
					maAlgo = "DoubleSMA"
				case 5:
					maAlgo = "DoubleRMA"
				case 6:
					maAlgo = "vwma"
				case 7:
					maAlgo = "SWMA"
				}
				for k := 0; k < 8; k++ {

					buy, sell := diamondTemplate(s, nil, e, k+2, true)
					strat := strategy.NewSimple(chart, buy, sell)
					de := fmt.Sprintf("%d | %s | %s | %d", i, srcAlgo, maAlgo, k-1)
					_, ii := strat.Results("side", true, "description", de)
					inst = append(inst, ii...)
					backtest++
					trades += strat.GetTrades().Len()
				}
			}
		}
	}
	sort.Sort(inst)
	hh := time.Since(nau)

	for _, i := range inst {
		fmt.Println(i.Description())
		i.Print()
	}
	fmt.Println(backtest, "Backtest with", trades, "Trades calculated in ", hh)
}

func mainSaphirIteration(chart *OHCLV) {
	open := Source(OPEN, chart)
	close := Source(CLOSE, chart)
	low := Source(LOW, chart)
	high := Source(HIGH, chart)
	volume := Source(VOLUME, chart)
	openclose := Source(OC2, chart)
	ha := HeikinAshi(chart)

	haHL2 := Source(HL2, ha)
	//ohclv := Source(OHCL4, chart)
	//hl2 := Source(HL2, chart)

	swma := func(s Series, l int) Series {
		_ = l
		return Swma(s)
	}

	maFuncs := []func(src Series, l1 int) Series{
		Ema,
		Sma,
		doubleEMA,
		doubleSMA,
		swma,
	}
	trades := 0
	backtest := 0
	nau := time.Now()

	var srcAlgo string
	var maAlgo string

	var inst strategy.Instances

	fmt.Println("Start Iterating")

	for i := 0; i < 7; i++ {
		Strategies := []Series{Rsi(openclose, 6+i), Ema(openclose, 6+i), Roc(openclose, 4+i), dai(open, high, low, close, volume, 7, 2+i), Swma(openclose)}
		Strategies2 := []Series{Rsi(haHL2, 6+i), Ema(haHL2, 6+i), Roc(haHL2, 4+i), dai(open, high, low, close, volume, 7, 2+i), Swma(haHL2)}
		for ss, s := range Strategies {
			switch ss {
			case 0:
				srcAlgo = "RSI"
			case 1:
				srcAlgo = "EMA"
			case 2:
				srcAlgo = "Roc"
			case 3:
				srcAlgo = "DAI"
			case 4:
				srcAlgo = "Swma"
			}
			for ee, e := range maFuncs {
				switch ee {
				case 0:
					maAlgo = "EMA"
				case 1:
					maAlgo = "SMA"
				case 2:
					maAlgo = "DoubleEma"
				case 3:
					maAlgo = "DoubleSMA"
				case 4:
					maAlgo = "SWMA"
				}
				for k := 0; k < 6; k++ {

					buy, sell := saphirTemplate(s, Strategies2[ss], volume, e, k+3)
					strat := strategy.NewSimple(chart, buy, sell)
					de := fmt.Sprintf("%d | %s | %s | %d", i, srcAlgo, maAlgo, k)
					_, ii := strat.Results("side", false, "description", de)
					inst = append(inst, ii...)
					backtest++
					trades += strat.GetTrades().Len()
				}
			}
		}
	}
	sort.Sort(inst)
	hh := time.Since(nau)

	for _, i := range inst {
		fmt.Println(i.Description())
		i.Print()
	}
	fmt.Println(backtest, "Backtest with", trades, "Trades calculated in ", hh)
}
