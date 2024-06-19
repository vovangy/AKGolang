package main

import "fmt"

func PrintNumbers(nums ...int) {
	for _, val := range nums {
		fmt.Println(val)
	}
}

func main() {
	PrintNumbers(1, 2, 3, 4, 5)
}
