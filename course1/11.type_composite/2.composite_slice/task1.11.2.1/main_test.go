package main

import "testing"

type testDataGetSubSlice struct {
	inputSlice []int
	start      int
	end        int
	expected   []int
}

func TestGetSubSlice(t *testing.T) {
	testCases := []testDataGetSubSlice{
		{inputSlice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, start: 2, end: 6, expected: []int{3, 4, 5, 6}},
	}

	for _, tc := range testCases {
		result := getSubSlice(tc.inputSlice, tc.start, tc.end)
		for i, val := range tc.expected {
			if result[i] != val {
				t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc.inputSlice, tc.expected, result)
			}
		}
	}
}
