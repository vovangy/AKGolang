package main

import "fmt"

func RemoveIDX(xs []int, idx int) []int {
	if idx >= len(xs) || idx < 0 {
		return []int{}
	}

	return append(append([]int{}, xs[:idx]...), xs[idx+1:]...)
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(RemoveIDX(a, 3))
}
