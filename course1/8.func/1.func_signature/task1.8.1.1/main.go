package main

import (
	"fmt"
	"math"
)

var CalculateCircleArea = func(radius float64) float64 {
	return math.Pi * radius * radius
}

var CalculateRectangleArea = func(width, height float64) float64 {
	return width * height
}

var CalculateTriangleArea = func(base, height float64) float64 {
	return base * height * 0.5
}

func main() {
	fmt.Println(CalculateCircleArea(3))
	fmt.Println(CalculateRectangleArea(3, 4))
	fmt.Println(CalculateTriangleArea(3, 4))
}
