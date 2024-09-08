package main

import (
	"testing"
)

func TestHashMap_SetAndGet(t *testing.T) {
	m := NewHashMap()
	m.Set("key", "value")
	if value, ok := m.Get("key"); !ok || value != "value" {
		t.Errorf("Expected 'value', got '%v'", value)
	}
}

func BenchmarkHashMap_CRC64(b *testing.B) {
	m := NewHashMap(WithHashCRC64())
	for i := 0; i < b.N; i++ {
		m.Set("key", "value")
		_, _ = m.Get("key")
	}
}

func BenchmarkHashMap_CRC32(b *testing.B) {
	m := NewHashMap(WithHashCRC32())
	for i := 0; i < b.N; i++ {
		m.Set("key", "value")
		_, _ = m.Get("key")
	}
}

func BenchmarkHashMap_CRC16(b *testing.B) {
	m := NewHashMap(WithHashCRC16())
	for i := 0; i < b.N; i++ {
		m.Set("key", "value")
		_, _ = m.Get("key")
	}
}

func BenchmarkHashMap_CRC8(b *testing.B) {
	m := NewHashMap(WithHashCRC8())
	for i := 0; i < b.N; i++ {
		m.Set("key", "value")
		_, _ = m.Get("key")
	}
}
