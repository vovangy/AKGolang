package main

import "testing"

type testDataConcat struct {
	input    []interface{}
	expected string
}

type testDataSum struct {
	input    []interface{}
	expected int
}

type testDataSumFloat struct {
	input    []interface{}
	expected float64
}

func TestOperateSum(t *testing.T) {
	testCasesInt := []testDataSum{
		{input: []interface{}{1, 1, 2, 3, 3, 4, 4, 5, 5}, expected: 28},
		{input: []interface{}{1, 2, 2, 3, 3, 4, 4}, expected: 19},
		{input: []interface{}{}, expected: 0},
	}
	testCasesFloat := []testDataSumFloat{
		{input: []interface{}{1.5, 1, 2, 3, 3, 4, 4, 5, 5.5}, expected: 29},
	}

	for _, tc := range testCasesInt {
		result := Operate(Sum, tc.input...)
		if result != tc.expected {
			t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc.input, tc.expected, result)
		}
	}
	for _, tc := range testCasesFloat {
		result := Operate(Sum, tc.input...)
		if result != tc.expected {
			t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc.input, tc.expected, result)
		}
	}
}

func TestOperateConcat(t *testing.T) {
	testCases := []testDataConcat{
		{input: []interface{}{"ZOLo", "u", "ZALA"}, expected: "ZOLouZALA"},
		{input: []interface{}{}, expected: ""},
	}

	for _, tc := range testCases {
		result := Operate(Concat, tc.input...)
		if result != tc.expected {
			t.Errorf("Unexpected result. Input: %v, Expected: %v, Got: %v", tc.input, tc.expected, result)
		}
	}
}
