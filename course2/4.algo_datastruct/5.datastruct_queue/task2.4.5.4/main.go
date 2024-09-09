package main

import (
	"fmt"
	"log"
)

var stack []int

func push(value int) {
	stack = append(stack, value)
}

func pop() int {
	if len(stack) == 0 {
		log.Fatal("Ошибка: Стек пустой!")
	}
	lastIndex := len(stack) - 1
	value := stack[lastIndex]
	stack = stack[:lastIndex]
	return value
}

func main() {
	push(5)
	push(3)
	result := pop() + pop()
	push(result)

	fmt.Println(stack[0])
}
