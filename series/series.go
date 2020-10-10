package series

import (
	"github.com/dawnkosmos/fastpine/exchange"
)

//Value is just the empty interface
type Value interface{}

//Series tries to achieve the pine script series data type
type Series interface {
	//Value ask for the Series value. 0 is the actual value, 1 the one before etc
	Value(int) float64
	//Resolution returns the resolution in seconds
	Resolution() int
	//Starttime is the point where the series has its first values
	Starttime() int64
	//Data returns the whole Data. Should only be used in init an indicator
	Data() *[]float64
	//UpdateGroup gets passed on from series to series.
	//So the update order is right
	UpdateGroup() *exchange.UpdateGroup
}

type Condition interface {
	ValueB(int) bool
	Starttime() int64
	DataB() *[]bool
	Resolution() int
	UpdateGroup() *exchange.UpdateGroup
}

//HELP FUNCS

func bigger(i, j int64) int64 {
	if i > j {
		return i
	}
	return j
}

func Average(f []float64) float64 {
	return Sum(f) / float64(len(f))
}

func Sum(f []float64) float64 {
	var avg float64 = 0
	for _, v := range f {
		avg += v
	}
	return avg
}

func Change(old, new float64) float64 {
	return new - old
}

func opExecute(op func(float64, float64) float64, f1 []float64, f2 []float64) []float64 {
	fOut := make([]float64, 0, len(f1))
	for i := 0; i < len(f1); i++ {
		fOut = append(fOut, op(f1[i], f2[i]))
	}
	return fOut
}

func mult(v1, v2 float64) float64 {
	return v1 * v2
}

/*	TODO
Bollinger Bands
Stdev
Lowest,Highest //Update
Stochastic
Atr
Heikin Ashi
Vwap
TSI
Sum
Roc
iff
Greater
Lower
Or
And
Functions

Buy/Sell
Backtesting
*/
