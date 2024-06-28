package main

import "fmt"

func InsertToStart(xs []int, x ...int) []int {
	var result []int
	result = append(result, x...)
	result = append(result, xs...)
	return result
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 9, 10}
	fmt.Println(InsertToStart(a, b...))
}
