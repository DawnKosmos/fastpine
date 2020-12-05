package series

import (
	"fmt"
	"strconv"

	"github.com/dawnkosmos/fastpine/exchange"
	"github.com/dawnkosmos/fastpine/helper"
)

/*
Hey you made it to this code.
Most algos are hardcoded to have a better performance.
Most Indicators are easy to calculate, though calculating them in a fast way can be hard and looks shit.
I also don't check for stupid inputs, so don't try to calculate the 1000 daily SMA and wonder
why you get no result. Or be funny and calculate negative lenght.
*/

//Value is just the empty interface
type Value interface{}

//Series is the Tradingview Series
type Series interface {
	StarttimeResolutionUpdategroup
	//Value ask for the Series value. 0 is the actual value, 1 the one before etc
	Value(int) float64
	//Data returns the whole Data. Should only be used in init an indicator
	Data() []float64
}

//Condition is the Series[bool] from tradingview, its needed to create tradingrules for strategies
type Condition interface {
	StarttimeResolutionUpdategroup
	//ValueB ask for the Series value. index = 0, is the actual value, 1 the one before etc.
	ValueB(int) bool
	//DataB return the saved data as an array. Should only be used to init an indicator. The last value represents the actual value
	DataB() []bool
}

/*
StartTimeResolutionUpdategroup is need for the interface to interact with each other,
starttime is needed to synchronize the indicators
the resolution is needed to  calculate the starttime
the UpdateGroup is needed so the indicators get updated the right way.
*/
type StarttimeResolutionUpdategroup interface {
	//Resolution returns the resolution in seconds
	Resolution() int
	//Starttime is the point where the series has its first values
	Starttime() int64
	//UpdateGroup gets passed on from series to series.
	//So the update order is right
	UpdateGroup() *exchange.UpdateGroup
}

//USR = updategroup Starttime Resolution
type USR struct {
	ug         *exchange.UpdateGroup
	starttime  int64
	resolution int
}

func (u *USR) UpdateGroup() *exchange.UpdateGroup {
	return u.ug
}

func (u *USR) Starttime() int64 {
	return u.starttime
}

func (u *USR) Resolution() int {
	return u.resolution
}

//HELP FUNCS

func bigger(i, j int64) int64 {
	if i > j {
		return i
	}
	return j
}

//Change returns the difference between a new and old value
func change(old, new float64) float64 {
	return new - old
}

//opExecutre executes an opperation on 2 float arrays which need to have the same lenght
func opExecute(op func(float64, float64) float64, f1 []float64, f2 []float64) []float64 {
	fOut := make([]float64, 0, len(f1))
	for i := 0; i < len(f1); i++ {
		fOut = append(fOut, op(f1[i], f2[i]))
	}
	return fOut
}

//mult is just v1 * v2
func mult(v1, v2 float64) float64 {
	return v1 * v2
}

func PrintSeries(i ...interface{}) {
	var SeriesData [][]float64
	var ConditionsData [][]bool

	for _, v := range i {
		switch f := v.(type) {
		case Series:
			SeriesData = append(SeriesData, f.Data())
		case Condition:
			ConditionsData = append(ConditionsData, f.DataB())
		}
	}

	var lenghts []int
	for _, v := range SeriesData {
		lenghts = append(lenghts, len(v))
	}

	for _, v := range ConditionsData {
		lenghts = append(lenghts, len(v))
	}

	smol := helper.Min(lenghts...)

	for i, v := range SeriesData {
		SeriesData[i] = v[len(v)-smol:]
	}

	for i, v := range ConditionsData {
		ConditionsData[i] = v[len(v)-smol:]
	}
	var s string

	for i := 0; i < smol; i++ {
		s = s + strconv.Itoa(i)
		for _, v := range SeriesData {
			s = s + " " + strconv.FormatFloat(v[i], 'f', 3, 64)
		}

		for _, v := range ConditionsData {
			s = s + " " + strconv.FormatBool(v[i])
		}

		fmt.Println(s)
		s = ""
	}

}

/*	TODO
Bollinger Bands
Stdev
Backtesting
*/
