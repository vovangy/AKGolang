package main

import "fmt"

func FindMaxAndMin(n ...int) (int, int) {
	curMax := 0
	curMin := 0
	if len(n) > 0 {
		curMax = n[0]
		curMin = n[0]
	}

	for _, val := range n {
		if curMax < val {
			curMax = val
		}
		if curMin > val {
			curMin = val
		}
	}

	return curMax, curMin
}

func main() {
	fmt.Println(FindMaxAndMin(1, 2, 3, 4, 5, 6))
}
