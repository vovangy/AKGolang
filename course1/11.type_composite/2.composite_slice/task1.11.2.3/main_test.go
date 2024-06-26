package main

import "testing"

type testDataFindSingleNumber struct {
	input    []int
	expected int
}

func TestFindSingleNumber(t *testing.T) {
	testCases := []testDataFindSingleNumber{
		{input: []int{1, 1, 2, 3, 3, 4, 4, 5, 5}, expected: 2},
		{input: []int{1, 2, 2, 3, 3, 4, 4}, expected: 1},
		{input: []int{}, expected: 0},
	}

	for _, tc := range testCases {
		result := findSingleNumber(tc.input)
		if result != tc.expected {
			t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc.input, tc.expected, result)
		}
	}
}
