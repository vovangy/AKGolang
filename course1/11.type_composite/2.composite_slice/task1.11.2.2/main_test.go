package main

import "testing"

type testDataMaxDifference struct {
	input    []int
	expected int
}

func TestSortDescInt(t *testing.T) {
	testCases := []testDataMaxDifference{
		{input: []int{1, 2, 3, 4, 5, 6, 7, 8}, expected: 7},
		{input: []int{8, 7, 6, 5, 4, 3, 2, 1}, expected: 7},
		{input: []int{1, 2}, expected: 1},
		{input: []int{1}, expected: 0},
		{input: []int{}, expected: 0},
	}

	for _, tc := range testCases {
		result := MaxDifference(tc.input)
		if result != tc.expected {
			t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc.input, tc.expected, result)
		}
	}
}
