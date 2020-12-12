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
	openclose := OC2(chart)
	//ohclv := Source(OHCL4, chart)
	//hl2 := Source(HL2, chart)

	maFuncs := []func(src Series, l1 int) Series{
		Ema,
		Sma,
		Rma,
		doubleEMA,
		doubleSMA,
	}
	trades := 0
	nau := time.Now()

	var srcAlgo string
	var maAlgo string

	var inst strategy.Instances

	//DaiCross RSICross
	fmt.Println("Start Iterating")

	for i := 0; i < 6; i++ {
		Strategies := []Series{Dai(open, high, low, close, volume, 4+i, 4), Ema(Stoch(openclose, high, low, 8+i), 4+i), Sma(openclose, 3+i), Sma(Rsi(openclose, 7+i), 2)}
		for ss, s := range Strategies {
			switch ss {
			case 0:
				srcAlgo = "dai"
			case 1:
				srcAlgo = "stoch"
			case 2:
				srcAlgo = "sma"
			case 3:
				srcAlgo = "rsi"
			}

			for ee, ma := range maFuncs {
				for k := 0; k < 8; k++ {
					buy, sell := CrossDoubleMA(s, ma, k+3, false)
					strat := strategy.NewSimple(chart, buy, sell)
					switch ee {
					case 0:
						maAlgo = "EMA"
					case 1:
						maAlgo = "SMA"
					case 2:
						maAlgo = "RMA"
					case 3:
						maAlgo = "DoubleEMA"
					case 4:
						maAlgo = "DoubleSMA"
					}
					de := fmt.Sprintf("%d | %s | %s | %d", i, srcAlgo, maAlgo, k+3)
					_, ii := strat.Results("side", true, "description", de)
					inst = append(inst, ii...)
					trades++

				}
			}

		}
	}
	sort.Sort(inst)

	for _, i := range inst {
		fmt.Println(i.Description())
		i.Print()
	}

	fmt.Println(trades, "Backtest calculated in ", time.Since(nau))
}

func mainDiamondIterations(chart *OHCLV) {
	open := Source(OPEN, chart)
	close := Source(CLOSE, chart)
	low := Source(LOW, chart)
	high := Source(HIGH, chart)
	volume := Source(VOLUME, chart)
	//openclose := OC2(chart)
	//ohclv := Source(OHCL4, chart)
	//hl2 := Source(HL2, chart)

	vwma := func(s Series, l int) Series {
		return Vwma(s, volume, l)
	}

	swma := func(s Series, l int) Series {
		_ = l
		return Swma(s)
	}

	maFuncs := []func(srctrue Series, l1 int) Series{
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

	fmt.Println("Start Iterating")

	loCo := longCon(0.0, 6, close, low)
	shCo := shortCon(0.0, 6, close, high)

	ha := HeikinAshi(chart)
	var inputName string

	input := []Series{open, close, HCL3(chart), OC2(ha), HL2(chart), OC2(chart)}
	for iii, ii := range input {
		switch iii {
		case 0:
			inputName = "open"
		case 1:
			inputName = "close"
		case 2:
			inputName = "HCL3"
		case 3:
			inputName = "HA OC2"
		case 4:
			inputName = "HL2"
		case 5:
			inputName = "OC2"
		}
		src := ii

		for i := 0; i < 7; i++ {
			Strategies := []Series{Rsi(src, 6+i), Ema(src, 6+i), Roc(src, 4+i), dai(open, high, low, close, volume, 7, 2+i), Swma(src), Sma(src, 4+i)}
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
				case 5:
					srcAlgo = "SMA"
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
						buy, sell := diamondTemplate(s, volume, e, k+2, false)
						strat := strategy.NewSimple(chart, And(buy, loCo), And(sell, shCo), "pyramiding", 3)
						de := fmt.Sprintf("%s | %d | %s | %s | %d", inputName, i, srcAlgo, maAlgo, k-1)
						_, ii := strat.Results("side", false, "description", de)
						inst = append(inst, ii...)
						backtest++

						trades += strat.GetTrades().Len()
						if ee == 7 {
							break
						}
					}
				}
			}
		}
	}
	sort.Sort(inst)
	hh := time.Since(nau)

	for _, i := range inst[len(inst)-50:] {
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
	openclose := OC2(chart)
	//ha := HeikinAshi(chart)

	haHL2 := HL2(chart)
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

	loCo := longCon(0.8, 6, close, low)
	shCo := shortCon(1.0, 6, close, high)

	for i := 0; i < 8; i++ {
		Strategies := []Series{Rsi(openclose, 6+i), Ema(openclose, 6+i), Roc(openclose, 4+i), dai(open, high, low, close, volume, 7, 2+i), Swma(openclose)}
		Strategies2 := []Series{Rsi(haHL2, 6+i), Ema(haHL2, 6+i), Roc(haHL2, 3+i), dai(open, high, low, close, volume, 7, 2+i), Swma(haHL2)}
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

					buy, sell := saphirTemplate(s, Strategies2[ss], nil, e, k+3)
					strat := strategy.NewSimple(chart, And(buy, loCo), And(sell, shCo), "pyramiding", 3)
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

	for _, i := range inst[len(inst)-40:] {
		fmt.Println(i.Description())
		i.Print()
	}
	fmt.Println(backtest, "Backtest with", trades, "Trades calculated in ", hh)
}
