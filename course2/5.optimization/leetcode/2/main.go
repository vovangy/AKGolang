package main

import (
	"fmt"
	"sort"
)

func sortStudentsByKthScore(score [][]int, k int) [][]int {
	sort.Slice(score, func(i, j int) bool {
		return score[i][k] > score[j][k]
	})
	return score
}

func main() {
	score := [][]int{
		{10, 20, 30},
		{40, 50, 60},
		{70, 80, 90},
	}
	k := 1
	sortedScores := sortStudentsByKthScore(score, k)

	for _, row := range sortedScores {
		fmt.Println(row)
	}
}
