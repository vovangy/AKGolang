package main

import (
	"fmt"
)

type CircuitRinger interface {
	Add(val int)
	Get() (int, bool)
}

type RingBuffer struct {
	buffer []int
	size   int
	start  int
	end    int
	isFull bool
}

func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		buffer: make([]int, size),
		size:   size,
		start:  0,
		end:    0,
		isFull: false,
	}
}

func (rb *RingBuffer) Add(val int) {
	rb.buffer[rb.end] = val
	if rb.isFull {
		rb.start = (rb.start + 1) % rb.size
	}
	rb.end = (rb.end + 1) % rb.size
	rb.isFull = rb.end == rb.start
}

func (rb *RingBuffer) Get() (int, bool) {
	if rb.start == rb.end && !rb.isFull {
		return 0, false
	}
	val := rb.buffer[rb.start]
	rb.start = (rb.start + 1) % rb.size
	rb.isFull = false
	return val, true
}

func main() {
	rb := NewRingBuffer(3)
	rb.Add(1)
	rb.Add(2)
	rb.Add(3)
	rb.Add(4)

	for val, ok := rb.Get(); ok; val, ok = rb.Get() {
		fmt.Println(val)
	}

	if _, ok := rb.Get(); !ok {
		fmt.Println("Buffer is empty")
	}
}
