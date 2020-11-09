package helper

import "sort"

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
func FloatAverage(f []float64) float64 {
	return FloatSum(f) / float64(len(f))
}

//FloatMedian returns the Median of a float array
func FloatMedian(f ...float64) float64 {
	sort.Float64s(f)

	median := len(f) / 2

	if median%2 == 1 {
		return f[median]
	}
	return (f[median-1] + f[median]) / 2
}

//FloatArrLowestLen return the lenght of the smallest inputed float array
func FloatArrLowestLen(f ...[]float64) int {
	l := len(f[0])
	for _, v := range f[1:] {
		if len(v) < l {
			l = len(v)
		}
	}
	return l

}

//FloatMax returns the highest Value
func FloatMax(f ...float64) float64 {
	max := f[0]
	for _, v := range f[1:] {
		if v > max {
			max = v
		}
	}

	return max
}

//FloatSum returns the sum of a float array
func FloatSum(f []float64) float64 {
	var avg float64 = 0
	for _, v := range f {
		avg += v
	}
	return avg
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
