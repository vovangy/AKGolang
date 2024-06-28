package main

import "fmt"

func Pop(xs []int) (int, []int) {
	value := xs[0]
	newStack := xs[1:]
	return value, newStack
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	val, newSlice := Pop(a)
	fmt.Printf("Значение: %d, Новый срез: %v", val, newSlice)
}
