package main

import "fmt"

func Shift(xs []int) (int, []int) {
	if len(xs) < 1 {
		return 0, xs
	}
	if len(xs) == 1 {
		return xs[0], xs
	}
	result := append(xs[len(xs)-1:], xs[:len(xs)-1]...) // Создаем слайс делая срез от последнего элемента и добавляя к нему срез от первого до последнего
	return result[1], result
}

func main() {
	xs := []int{1, 2, 3, 4, 5}
	firstElement, shiftedSlice := Shift(xs)
	fmt.Println(firstElement)
	fmt.Println(shiftedSlice)
}
