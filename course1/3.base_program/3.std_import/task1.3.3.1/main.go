package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	if x < 0 {
		return 0
	}

	return math.Sqrt(x)
}

func main() {
	fmt.Println(Sqrt(4))
}
