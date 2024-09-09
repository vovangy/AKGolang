package main

import (
	"testing"
)

func TestFactorialRecursive(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 1},
		{5, 120},
		{10, 3628800},
	}

	for _, test := range tests {
		result := factorialRecursive(test.input)
		if result != test.expected {
			t.Errorf("FactorialRecursive(%d) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

func TestFactorialIterative(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 1},
		{5, 120},
		{10, 3628800},
	}

	for _, test := range tests {
		result := factorialIterative(test.input)
		if result != test.expected {
			t.Errorf("FactorialIterative(%d) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

func TestCompareWhichFactorialIsFaster(t *testing.T) {
	results := compareWhichFactorialIsFaster()

	for input, isRecursiveFaster := range results {

		if isRecursiveFaster {
			t.Errorf("%s: expected iterative to be faster", input)
		}

	}
}
