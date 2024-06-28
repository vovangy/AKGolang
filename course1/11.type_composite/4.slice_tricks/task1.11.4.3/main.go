package main

import "fmt"

func RemoveExtraMemory(xs []int) []int {
	result := make([]int, len(xs), len(xs))
	copy(result, xs)
	return result
}

func main() {
	a := make([]int, 3, 6)
	b := RemoveExtraMemory(a)
	fmt.Println(b)
}
