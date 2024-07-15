package main

import (
	"fmt"
	"math"
)

func Sin(x float64) float64 {
	return math.Sin(x)
}
func Cos(x float64) float64 {
	return math.Cos(x)
}

func main() {
	fmt.Println(Cos(60)*Cos(60) + Sin(60)*Sin(60))
}
