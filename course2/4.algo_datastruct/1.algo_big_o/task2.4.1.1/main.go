package main

import (
	"fmt"
	"time"
)

func factorialRecursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorialRecursive(n-1)
}

func factorialIterative(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func compareWhichFactorialIsFaster() map[string]bool {
	results := make(map[string]bool)

	for i := 10; i <= 100; i += 10 {
		startRecursive := time.Now()
		factorialRecursive(i)
		durationRecursive := time.Since(startRecursive)

		startIterative := time.Now()
		factorialIterative(i)
		durationIterative := time.Since(startIterative)

		results[fmt.Sprintf("Input %d", i)] = durationRecursive < durationIterative
	}

	return results
}
