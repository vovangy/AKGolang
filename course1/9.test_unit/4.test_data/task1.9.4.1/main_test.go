package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func generateSlice(size int) []float64 {
	var testData []float64
	for i := 0; i < size; i++ {
		testData = append(testData, rand.Float64())
	}
	return testData
}

func TestAverage(t *testing.T) {
	testData := generateSlice(9)

	myAverage := average(testData)
	testAverage := average(testData)

	reflect.DeepEqual(myAverage, testAverage)
}
