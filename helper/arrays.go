package helper

import (
	"math"
	"sort"
)

//FLOAT

//PositionOfLowestValue return the Position of the lowest Value of a Float Array
func FloatPositionOfLowestValue(f []float64) int {
	low := f[0]
	lowPosition := 0

	for i := 1; i < len(f); i++ {
		if f[i] <= low {

			lowPosition = i
			low = f[i]
		}
	}
	return lowPosition
}

//PositionOfHighestValue return the Position of the highest Value of a Float Array
func FloatPositionOfHighestValue(f []float64) int {
	high := f[0]
	highPosition := 0

	for i := 1; i < len(f); i++ {
		if f[i] >= high {

			highPosition = i
			high = f[i]
		}
	}
	return highPosition
}

//FloatAverage returns the mean of a float array
func FloatAverage(arr []float64) float64 {
	return FloatSum(arr) / float64(len(arr))
}

//FloatMedian returns the Median of a float array
func FloatMedian(arr ...float64) float64 {
	sort.Float64s(arr)

	median := len(arr) / 2

	if median%2 == 1 {
		return arr[median]
	}
	return (arr[median-1] + arr[median]) / 2
}

//FloatArrLowestLen return the lenght of the smallest inputed float array
func FloatArrLowestLen(arr ...[]float64) int {
	l := len(arr[0])
	for _, v := range arr[1:] {
		if len(v) < l {
			l = len(v)
		}
	}
	return l

}

//FloatMax returns the highest Value
func FloatMax(arr ...float64) float64 {
	max := arr[0]
	for _, v := range arr[1:] {
		if v > max {
			max = v
		}
	}

	return max
}

//FloatSum returns the sum of a float array
func FloatSum(arr []float64) float64 {
	var avg float64 = 0
	for _, v := range arr {
		avg += v
	}
	return avg
}

//FloatStdev calculates the standard derivation of an array
func FloatStdev(arr []float64) float64 {
	N := float64(len(arr) - 1)
	mean := FloatAverage(arr)
	fn := func(v float64) float64 {
		return math.Pow(v-mean, 2.0)
	}
	sArr := FloatOperator(arr, fn)

	return math.Sqrt(FloatSum(sArr) / N)
}

//FloatOperator uses a function on each individual value of the float array
func FloatOperator(arr []float64, fn func(float64) float64) []float64 {
	out := make([]float64, 0, len(arr))
	for _, v := range arr {
		out = append(out, fn(v))
	}
	return out
}

//INTEGER

//Min returns the smallest value of the inputed parameters
func Min(i ...int) int {
	lowest := i[0]
	for _, v := range i {
		if v < lowest {
			lowest = v
		}
	}

	return lowest
}

//Max returns the biggest value of the inputed parameters
func Max(i ...int) int {
	h := i[0]
	for _, v := range i {
		if v > h {
			h = v
		}
	}

	return h
}

//IntSum returns the sum of a int array
func IntSum(f []int) int {
	var avg int = 0
	for _, v := range f {
		avg += v
	}
	return avg
}
