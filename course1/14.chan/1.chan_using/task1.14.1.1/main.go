package main

import (
	"fmt"
	"time"
)

func generateChan(n int) chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func mergeChan(mergeTo chan int, from ...chan int) {
	go func() {
		for _, ch := range from {
			for i := range ch {
				mergeTo <- i
			}
		}
		close(mergeTo)
	}()
}

func mergeChan2(chans ...chan int) chan int {
	mergeCh := make(chan int)
	go func() {
		for _, ch := range chans {
			for i := range ch {
				mergeCh <- i
			}
		}
		close(mergeCh)
	}()
	return mergeCh
}

func main() {
	a := generateChan(10)
	b := generateChan(8)
	c := generateChan(6)
	d := make(chan int)

	e := mergeChan2(a, b, c)

	mergeChan(d, a, b, c)

	//for i := range d {
	//	fmt.Println(i)
	//}

	for i := range e {
		fmt.Println(i)
	}

	time.Sleep(time.Second)
}
