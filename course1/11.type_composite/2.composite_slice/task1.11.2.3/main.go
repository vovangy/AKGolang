package main

func bitwiseXOR(n, res int) int {
	return n ^ res
}

func findSingleNumber(numbers []int) int {
	result := 0
	for _, val := range numbers {
		result = bitwiseXOR(result, val)
	}

	return result
}
