package main

import (
	"fmt"
	"sort"
)

func isArithmetic(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sort.Ints(nums)

	diff := nums[1] - nums[0]

	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] != diff {
			return false
		}
	}
	return true
}

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	m := len(l)
	results := make([]bool, m)

	for i := 0; i < m; i++ {
		subarray := nums[l[i] : r[i]+1]

		results[i] = isArithmetic(subarray)
	}

	return results
}

func main() {
	nums := []int{4, 6, 5, 9, 3, 7}
	l := []int{0, 0, 2}
	r := []int{2, 3, 5}

	result := checkArithmeticSubarrays(nums, l, r)
	fmt.Println(result)
}
