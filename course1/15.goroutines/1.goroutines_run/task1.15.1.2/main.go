package main

import (
	"fmt"
	"time"
)

func NotifyEvery(ticker *time.Ticker, d time.Duration, message string) <-chan string {
	data := make(chan string)

	go func() {
		defer close(data)

		timeout := time.After(d)
		for {
			select {
			case <-timeout:
				return
			case <-ticker.C:
				data <- message
			}
		}
	}()

	return data
}

func main() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	duration := 5 * time.Second
	message := "Таймер сработал"
	data := NotifyEvery(ticker, duration, message)

	for v := range data {
		fmt.Println(v)
	}

	fmt.Println("Программа завершена")
}
