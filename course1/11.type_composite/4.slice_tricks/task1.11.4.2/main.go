package main

import "fmt"

func InsertAfterIDX(xs []int, idx int, x ...int) []int {
	result := append(append(xs[:idx+1], x...), xs[idx+1:]...)
	return result
}

func main() {
	xs := []int{1, 2, 3, 4, 5}
	result := InsertAfterIDX(xs, 2, 6, 7, 8)
	fmt.Println(result)
}
