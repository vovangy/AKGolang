package main

import (
	"strings"
	"testing"
)

func TestGenerateActivationKeyLength(t *testing.T) {
	key := generateActivationKey()
	expectedLength := 19
	if len(key) != expectedLength {
		t.Errorf("Expected length %d, but got %d", expectedLength, len(key))
	}
}

func TestGenerateActivationKeyFormat(t *testing.T) {
	key := generateActivationKey()
	parts := strings.Split(key, "-")
	if len(parts) != 4 {
		t.Errorf("Expected 4 parts separated by '-', but got %d parts", len(parts))
	}

	for _, part := range parts {
		if len(part) != 4 {
			t.Errorf("Each part should be of length 4, but got length %d", len(part))
		}
	}
}

func TestGenerateActivationKeyRandomness(t *testing.T) {
	key1 := generateActivationKey()
	key2 := generateActivationKey()
	if key1 == key2 {
		t.Errorf("Expected different keys for two calls, but got the same: %s", key1)
	}
}
