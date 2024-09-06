package main

import (
	"fmt"
	"time"
)

func main() {
	timeFormat := "15:04:05"
	dateFormat := "2006-01-02"

	for {
		now := time.Now()

		currentTime := now.Format(timeFormat)
		currentDate := now.Format(dateFormat)

		fmt.Print("\033[H\033[2J")

		fmt.Printf("Текущее время: %s\n", currentTime)
		fmt.Printf("Текущая дата: %s\n", currentDate)

		time.Sleep(1 * time.Second)
	}
}
