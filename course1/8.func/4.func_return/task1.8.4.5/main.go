package main

import (
	"fmt"
	"math"
)

func CalculatePercentageChange(initialValue, finalValue float64) float64 {
	if initialValue == 0 {
		return 0
	}
	percentageChange := ((finalValue - initialValue) / initialValue) * 100
	roundedPercentageChange := math.Round(percentageChange*100) / 100

	return roundedPercentageChange
}

func main() {
	fmt.Println(CalculatePercentageChange(100, 120.223))
}
