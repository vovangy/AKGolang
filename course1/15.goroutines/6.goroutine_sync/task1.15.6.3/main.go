package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
	fmt.Println(c.value)

	return c.value
}

func concurrentSafeCounter() int {
	counter := Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()

	return counter.value
}

func main() {
	_ = concurrentSafeCounter()
	//fmt.Println(result)
}
