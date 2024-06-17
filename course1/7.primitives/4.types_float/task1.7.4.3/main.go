package main

import (
	"fmt"
	"math"
)

func RoundToDecimalPlaces(value float64, decimalPlaces int) float64 {
	scale := math.Pow(10, float64(decimalPlaces))
	return math.Round(value*scale) / scale
}

func CompareRoundedValues(a, b float64, decimalPlaces int) (isEqual bool, difference float64) {
	roundedA := RoundToDecimalPlaces(a, decimalPlaces)
	roundedB := RoundToDecimalPlaces(b, decimalPlaces)

	isEqual = roundedA == roundedB
	difference = roundedA - roundedB

	return isEqual, difference
}

func main() {
	a := 3.1415926535
	b := 3.1415926536
	decimalPlaces := 10

	isEqual, difference := CompareRoundedValues(a, b, decimalPlaces)
	fmt.Printf("Values are equal: %v, Difference: %f\n", isEqual, difference)
}
