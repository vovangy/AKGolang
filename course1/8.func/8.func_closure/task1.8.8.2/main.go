package main

import "fmt"

func multiplier(factor float64) func(float64) float64 {
	return func(num float64) float64 {
		return num * factor
	}
}

func main() {
	m := multiplier(2.5)
	result := m(10)
	fmt.Println(result)
}
