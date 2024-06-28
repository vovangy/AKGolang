package main

import "fmt"

func FilterDividers(xs []int, divider int) []int {
	var result []int
	for _, val := range xs {
		if val%divider == 0 {
			result = append(result, val)
		}
	}
	return result
}

func main() {
	fmt.Println(FilterDividers([]int{1, 2, 3, 4, 5, 6, 7}, 2))
}
