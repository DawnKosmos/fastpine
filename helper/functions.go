package helper

import (
	"math"
)

//Package helper contains helper functions

//Change returns the difference between a new and old value
func Change(old, new float64) float64 {
	return new - old
}

//Int64Max return the bigger Value of 2 int64 numbers
func Int64Max(i, j int64) int64 {
	if i > j {
		return i
	} else {
		return j
	}
}

//Xor is the boolean operator xor
func Xor(x, y bool) bool {
	return (x || y) && !(x && y)
}

func FloatAbs(f float64) float64 {
	return math.Abs(f)
}
