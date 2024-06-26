package main

func MaxDifference(numbers []int) int {
	if len(numbers) < 2 {
		return 0
	}

	curMax := numbers[0]
	curMin := numbers[0]

	for _, val := range numbers {
		if curMax < val {
			curMax = val
		}
		if curMin > val {
			curMin = val
		}
	}

	return curMax - curMin
}
