package main

import (
	"fmt"
	"os"
	"time"
)

const (
	one = iota + 1
	two
	three
	four
	five
)

var ch = make(chan int)

func main() {
	numbers := []int{one, two, three, four, five}
	storeNumbers(numbers)
	print(ch)
}

func print(ch chan int) {
	if len(os.Getenv("DEBUG")) == 0 {
		fmt.Println("TUT")
		return
	}
	go func() {
		time.Sleep(1 * time.Second)
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}
}

func storeNumbers(numbers []int) {
	for _, num := range numbers {
		go func(num int) {
			write(num)
		}(num)
	}
}

func write(n int) {
	ch <- n
}
