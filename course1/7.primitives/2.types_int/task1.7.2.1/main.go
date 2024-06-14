package main

import "fmt"

func calculate(a int, b int) (int, int, int, int, int) {
	var sum int = a + b
	var difference int = a - b
	var product int = a * b
	var quotient int
	var remainder int
	if b != 0 {
		quotient = a / b
		remainder = a % b
	}
	return sum, difference, product, quotient, remainder
}

func main() {
	var sum int
	var difference int
	var product int
	var quotient int
	var remainder int

	sum, difference, product, quotient, remainder = calculate(10, 3)
	fmt.Printf("sum = %d difference = %d product = %d quotient = %d remainder = %d", sum, difference, product, quotient, remainder)
}
