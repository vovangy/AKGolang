package main

import (
	"math/rand"
	"testing"
	"time"
)

func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func TestInsertionSort(t *testing.T) {
	data := []int{34, 25, 12, 22, 11, 90}
	insertionSort(data)
	if !isSorted(data) {
		t.Errorf("InsertionSort failed. Sorted data: %v", data)
	}
}

func TestSelectionSort(t *testing.T) {
	data := []int{34, 25, 12, 22, 11, 90}
	selectionSort(data)
	if !isSorted(data) {
		t.Errorf("SelectionSort failed. Sorted data: %v", data)
	}
}

func TestQuicksort(t *testing.T) {
	data := []int{34, 25, 12, 22, 11, 90}
	sortedData := quicksort(data)
	if !isSorted(sortedData) {
		t.Errorf("Quicksort failed. Sorted data: %v", sortedData)
	}
}

func TestMergeSort(t *testing.T) {
	data := []int{34, 25, 12, 22, 11, 90}
	sortedData := mergeSort(data)
	if !isSorted(sortedData) {
		t.Errorf("MergeSort failed. Sorted data: %v", sortedData)
	}
}

func TestGeneralSort(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	tests := []struct {
		name string
		data []int
	}{
		{"Small array", []int{34, 25, 12, 22, 11, 90}},
		{"Medium array", []int{34, 25, 12, 22, 11, 90, 44, 33, 55, 77, 88, 99}},
		{"Large array", []int{34, 25, 12, 22, 11, 90, 44, 33, 55, 77, 88, 99, 34, 25, 12}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := make([]int, len(tt.data))
			copy(data, tt.data)
			generalSort(data)
			if !isSorted(data) {
				t.Errorf("GeneralSort failed. Sorted data: %v", data)
			}
		})
	}
}
