package main

import (
	"fmt"
	"hash/crc32"
	"testing"
)

func hashFunc(key string) int {
	return int(crc32.ChecksumIEEE([]byte(key)))
}

func BenchmarkHashMapArray(b *testing.B) {
	h := NewHashMapArray(1000, hashFunc)
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%d", i)
		h.Set(key, i)
		h.Get(key)
	}
}

func BenchmarkHashMapList(b *testing.B) {
	h := NewHashMapList(1000, hashFunc)
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%d", i)
		h.Set(key, i)
		h.Get(key)
	}
}

func TestHashMapArray(t *testing.T) {
	h := NewHashMapArray(16, hashFunc)

	h.Set("key1", "value1")
	h.Set("key2", "value2")

	if v, ok := h.Get("key1"); !ok || v != "value1" {
		t.Errorf("expected key1 to return value1")
	}

	if v, ok := h.Get("key2"); !ok || v != "value2" {
		t.Errorf("expected key2 to return value2")
	}
}

func TestHashMapList(t *testing.T) {
	h := NewHashMapList(16, hashFunc)

	h.Set("key1", "value1")
	h.Set("key2", "value2")

	if v, ok := h.Get("key1"); !ok || v != "value1" {
		t.Errorf("expected key1 to return value1")
	}

	if v, ok := h.Get("key2"); !ok || v != "value2" {
		t.Errorf("expected key2 to return value2")
	}
}
