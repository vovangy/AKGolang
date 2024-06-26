package main

import "testing"

type testDataAppendInt struct {
	inputSlice   []int
	inputNumbers []int
	expected     []int
}

func TestAppendInt(t *testing.T) {
	testCases := []testDataAppendInt{
		{inputSlice: []int{1, 2}, inputNumbers: []int{3, 4}, expected: []int{1, 2, 3, 4}},
	}

	for _, tc := range testCases {
		result := appendInt(tc.inputSlice, tc.inputNumbers...)
		for i, val := range tc.expected {
			if result[i] != val {
				t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc.inputSlice, tc.expected, result)
			}
		}
	}
}
