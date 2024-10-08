package main

import (
	"fmt"
)

func main() {
	a, b := 8, 13
	fmt.Println(*testDefer(&a, &b))
}

func testDefer(a, b *int) *int {
	c := 0

	defer func() {
		c = multiply(*a, *b)
	}()

	return &c
}

func sum(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}
