package main

import "testing"

func TestFactorial(t *testing.T) {
	result := Factorial(0)
	if result != 1 {
		t.Errorf("Factorial(0) = %d; want 1", result)
	}
	result = Factorial(1)
	if result != 1 {
		t.Errorf("Factorial(1) = %d; want 1", result)
	}
	result = Factorial(5)
	if result != 120 {
		t.Errorf("Factorial(5) = %d; want 120", result)
	}
}
