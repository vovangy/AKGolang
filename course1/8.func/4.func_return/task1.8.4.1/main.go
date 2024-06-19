package main

import "fmt"

func DivideAndRemainder(a, b int) (int, int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Деление на ноль")
		}
	}()

	return a / b, a % b
}

func main() {
	Divide, Remainder := DivideAndRemainder(4, 2)
	fmt.Printf("Частное: %d, Остаток: %d", Divide, Remainder)
}
