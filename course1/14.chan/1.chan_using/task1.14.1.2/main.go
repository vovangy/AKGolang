package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func generateData(n int) chan int {
	defer wg.Done()
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			ch <- i
		}
	}()
	return ch
}

func main() {
	wg.Add(1)
	data := generateData(10)
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		close(data)
	}()

	for num := range data {
		fmt.Println(num)
	}

	wg.Wait()
}
