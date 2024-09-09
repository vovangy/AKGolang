package main

import (
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name     string
		arr1     []User
		arr2     []User
		expected []User
	}{
		{
			name: "Both arrays are non-empty",
			arr1: []User{
				{ID: 1, Name: "Alice", Age: 30},
				{ID: 3, Name: "Bob", Age: 25},
				{ID: 5, Name: "Charlie", Age: 35},
			},
			arr2: []User{
				{ID: 2, Name: "David", Age: 40},
				{ID: 4, Name: "Eve", Age: 22},
				{ID: 6, Name: "Frank", Age: 45},
			},
			expected: []User{
				{ID: 1, Name: "Alice", Age: 30},
				{ID: 2, Name: "David", Age: 40},
				{ID: 3, Name: "Bob", Age: 25},
				{ID: 4, Name: "Eve", Age: 22},
				{ID: 5, Name: "Charlie", Age: 35},
				{ID: 6, Name: "Frank", Age: 45},
			},
		},
		{
			name: "First array is empty",
			arr1: []User{},
			arr2: []User{
				{ID: 2, Name: "David", Age: 40},
				{ID: 4, Name: "Eve", Age: 22},
			},
			expected: []User{
				{ID: 2, Name: "David", Age: 40},
				{ID: 4, Name: "Eve", Age: 22},
			},
		},
		{
			name: "Second array is empty",
			arr1: []User{
				{ID: 1, Name: "Alice", Age: 30},
				{ID: 3, Name: "Bob", Age: 25},
			},
			arr2: []User{},
			expected: []User{
				{ID: 1, Name: "Alice", Age: 30},
				{ID: 3, Name: "Bob", Age: 25},
			},
		},
		{
			name:     "Both arrays are empty",
			arr1:     []User{},
			arr2:     []User{},
			expected: []User{},
		},
		{
			name: "Identical arrays",
			arr1: []User{
				{ID: 1, Name: "Alice", Age: 30},
				{ID: 3, Name: "Bob", Age: 25},
			},
			arr2: []User{
				{ID: 1, Name: "Alice", Age: 30},
				{ID: 3, Name: "Bob", Age: 25},
			},
			expected: []User{
				{ID: 1, Name: "Alice", Age: 30},
				{ID: 1, Name: "Alice", Age: 30},
				{ID: 3, Name: "Bob", Age: 25},
				{ID: 3, Name: "Bob", Age: 25},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Merge(tt.arr1, tt.arr2)
			if len(result) != len(tt.expected) {
				t.Errorf("unexpected length: got %d, want %d", len(result), len(tt.expected))
				return
			}
			for i, user := range result {
				if user != tt.expected[i] {
					t.Errorf("unexpected user at index %d: got %+v, want %+v", i, user, tt.expected[i])
				}
			}
		})
	}
}
