package strategy

func Product(arr []float64) float64 {
	result := 1.0
	for _, v := range arr {
		result = result * (1 + v)
	}
	return result
}

func Winrate(arr []float64) float64 {
	amount := len(arr)
	var winners int

	for _, v := range arr {
		if v > 0 {
			winners++
		}
	}

	return float64(winners) / float64(amount) * 100
}
