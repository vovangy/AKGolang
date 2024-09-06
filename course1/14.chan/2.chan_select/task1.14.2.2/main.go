package main

import (
	"fmt"
	"time"
)

func timeout(timeout time.Duration) func() bool {
	timeoutCh := make(chan struct{})

	go func() {
		time.Sleep(timeout)
		close(timeoutCh)
	}()

	return func() bool {
		select {
		case <-timeoutCh:
			return false
		default:
			return true
		}
	}
}

func main() {
	timeoutFunc := timeout(3 * time.Second)
	since := time.NewTimer(4 * time.Second)

	for {
		select {
		case <-since.C:
			fmt.Println("Функция не выполнена вовремя")
			return
		default:
			if timeoutFunc() {
				fmt.Println("Функция выполнена вовремя")
				return
			}
		}
	}
}
