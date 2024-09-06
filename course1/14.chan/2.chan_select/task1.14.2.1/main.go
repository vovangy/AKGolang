package main

import "fmt"

func trySend(ch chan int, v int) bool {
	select {
	case ch <- v:
		return true
	default:
		return false
	}
}

func main() {
	ch := make(chan int, 1)

	fmt.Println(trySend(ch, 42))
	fmt.Println(trySend(ch, 43))

	close(ch)
}
