package main

import "fmt"

func createCounter() func() int {
	counter := 0
	return func() int {
		counter++
		return counter
	}
}

func main() {
	counter := createCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}
