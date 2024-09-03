package main

import (
	"reflect"
	"testing"
)

func TestCountWordsInText(t *testing.T) {
	tests := []struct {
		txt      string
		words    []string
		expected map[string]int
	}{
		{
			txt:      "Lorem ipsum dolor sit amet , consectetur adipiscing elit. Sit amet .",
			words:    []string{"sit", "amet", "lorem"},
			expected: map[string]int{"sit": 2, "amet": 2, "lorem": 1},
		},
		{
			txt:      "Hello world ! This is a test . Hello again.",
			words:    []string{"hello", "world", "test"},
			expected: map[string]int{"hello": 2, "world": 1, "test": 1},
		},
		{
			txt:      "Go is an open-source programming language that makes it easy to build simple, reliable, and efficient software.",
			words:    []string{"go", "easy", "efficient"},
			expected: map[string]int{"go": 1, "easy": 1, "efficient": 1},
		},
		{
			txt:      "No matching words here.",
			words:    []string{"absent", "missing", "none"},
			expected: map[string]int{"absent": 0, "missing": 0, "none": 0},
		},
	}

	for _, test := range tests {
		result := CountWordsInText(test.txt, test.words)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For text: %q and words: %v, expected %v, but got %v", test.txt, test.words, test.expected, result)
		}
	}
}
