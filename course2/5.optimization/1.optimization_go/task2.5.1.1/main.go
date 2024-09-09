package main

type HashMaper interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

// Вспомогательная структура для хранения пары ключ-значение
type arrayEntry struct {
	key   string
	value interface{}
}

type HashMapArray struct {
	buckets  [][]arrayEntry
	hashFunc func(string) int
}

func NewHashMapArray(size int, hashFunc func(string) int) *HashMapArray {
	return &HashMapArray{
		buckets:  make([][]arrayEntry, size),
		hashFunc: hashFunc,
	}
}

func (h *HashMapArray) Set(key string, value interface{}) {
	index := h.hashFunc(key) % len(h.buckets)
	h.buckets[index] = append(h.buckets[index], arrayEntry{key, value})
}

func (h *HashMapArray) Get(key string) (interface{}, bool) {
	index := h.hashFunc(key) % len(h.buckets)
	for _, entry := range h.buckets[index] {
		if entry.key == key {
			return entry.value, true
		}
	}
	return nil, false
}

// Вспомогательная структура для элемента списка
type listNode struct {
	key   string
	value interface{}
	next  *listNode
}

type HashMapList struct {
	buckets  []*listNode // Массив указателей на начало списков
	hashFunc func(string) int
}

func NewHashMapList(size int, hashFunc func(string) int) *HashMapList {
	return &HashMapList{
		buckets:  make([]*listNode, size),
		hashFunc: hashFunc,
	}
}

func (h *HashMapList) Set(key string, value interface{}) {
	index := h.hashFunc(key) % len(h.buckets)
	newNode := &listNode{key, value, nil}

	if h.buckets[index] == nil {
		h.buckets[index] = newNode
	} else {
		current := h.buckets[index]
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (h *HashMapList) Get(key string) (interface{}, bool) {
	index := h.hashFunc(key) % len(h.buckets)
	current := h.buckets[index]

	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}
	return nil, false
}
