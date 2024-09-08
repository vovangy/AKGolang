package main

import (
	"fmt"
)

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	inDegrees := make([]int, n)

	for _, edge := range edges {
		to := edge[1]
		inDegrees[to]++
	}

	var result []int
	for i, degree := range inDegrees {
		if degree == 0 {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	n := 6
	edges := [][]int{
		{0, 1},
		{0, 2},
		{2, 3},
		{4, 3},
		{4, 5},
	}

	result := findSmallestSetOfVertices(n, edges)
	fmt.Println("Smallest set of vertices:", result)
}
