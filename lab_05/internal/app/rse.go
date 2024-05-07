package app

import (
	"math"
)

func calcStdev(data []int64, avr float64) float64 {
	sum := float64(0)
	for _, val := range data {
		diff := float64(val) - avr
		sum += diff * diff
	}
	return math.Sqrt(sum / float64(len(data)-1))
}

func calcRSE(length int64, avr float64, stdev float64) float64 {
	return stdev / math.Sqrt(float64(length)) / avr * 100
}
