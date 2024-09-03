package main

import (
	"fmt"
	"time"
)

type sema struct {
	ch chan struct{}
}

func New(n int) sema {
	s := sema{ch: make(chan struct{}, n)}
	for i := 0; i < n; i++ {
		s.ch <- struct{}{}
	}
	return s
}

func (s sema) Inc(k int) {
	for i := 0; i < k; i++ {
		s.ch <- struct{}{}
	}
}

func (s sema) Dec(k int) {
	for i := 0; i < k; i++ {
		<-s.ch
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	n := len(numbers)
	sem := New(n)

	//var wg sync.WaitGroup
	//wg.Add(n)

	for _, num := range numbers {
		go func(n int) {
			//defer wg.Done()
			fmt.Println(n)
			sem.Inc(1)
		}(num)
	}

	sem.Dec(n)

	time.Sleep(time.Second)

	//wg.Wait()
}
