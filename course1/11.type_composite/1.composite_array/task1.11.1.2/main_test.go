package main

import "testing"

type testDataSortDescInt struct {
	input    [8]int
	expected [8]int
}

func TestSortDescInt(t *testing.T) {
	testCases := []testDataSortDescInt{
		{input: [8]int{1, 2, 3, 4, 5, 6, 7, 8}, expected: [8]int{8, 7, 6, 5, 4, 3, 2, 1}},
	}

	for _, tc := range testCases {
		result := sortDescInt(tc.input)
		for i, val := range tc.expected {
			if result[i] != val {
				t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc.input, tc.expected, result)
			}
		}
	}
}

type testDataSortAscInt struct {
	input    [8]int
	expected [8]int
}

func TestSortAscInt(t *testing.T) {
	testCases := []testDataSortAscInt{
		{input: [8]int{8, 7, 6, 5, 4, 3, 2, 1}, expected: [8]int{1, 2, 3, 4, 5, 6, 7, 8}},
	}

	for _, tc := range testCases {
		result := sortAscInt(tc.input)
		for i, val := range tc.expected {
			if result[i] != val {
				t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc.input, tc.expected, result)
			}
		}
	}
}

type testDataSortAscFloat struct {
	input    [8]float64
	expected [8]float64
}

func TestSortAscFloat(t *testing.T) {
	testCases := []testDataSortAscFloat{
		{input: [8]float64{8.5, 7.5, 6.5, 5.5, 4.5, 3.5, 2.5, 1.5}, expected: [8]float64{1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5}},
	}

	for _, tc := range testCases {
		result := sortAscFloat(tc.input)
		for i, val := range tc.expected {
			if result[i] != val {
				t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc.input, tc.expected, result)
			}
		}
	}
}

type testDataSortDescFloat struct {
	input    [8]float64
	expected [8]float64
}

func TestSortDescFloat(t *testing.T) {
	testCases := []testDataSortDescFloat{
		{input: [8]float64{1.5, 2.5, 3.5, 4.5, 5.5, 6.5, 7.5, 8.5}, expected: [8]float64{8.5, 7.5, 6.5, 5.5, 4.5, 3.5, 2.5, 1.5}},
	}

	for _, tc := range testCases {
		result := sortDescFloat(tc.input)
		for i, val := range tc.expected {
			if result[i] != val {
				t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc.input, tc.expected, result)
			}
		}
	}
}
