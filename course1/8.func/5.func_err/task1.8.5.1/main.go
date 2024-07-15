package main

import (
	"fmt"
	"math"
	"strconv"
)

func CalculatePercentageChange(initialValue, finalValue string) (float64, error) {
	initialValueFloat, err := strconv.ParseFloat(initialValue, 64)
	if err != nil {
		return 0, err
	}

	finalValueFloat, err := strconv.ParseFloat(finalValue, 64)
	if err != nil {
		return 0, err
	}

	if initialValueFloat <= 0 {
		return 0, fmt.Errorf("Initial value cannot be zero or negative")
	}
	percentageChange := ((finalValueFloat - initialValueFloat) / initialValueFloat) * 100
	roundedPercentageChange := math.Round(percentageChange*100) / 100

	return roundedPercentageChange, nil
}

func main() {
	fmt.Println(CalculatePercentageChange("100", "120.223"))
}
