package main

import (
	"fmt"
)

func getVariableType(variable interface{}) string {
	return fmt.Sprintf("%T", variable)
}

func main() {
	num := 10
	str := "Hello"
	floating := 3.14
	boolean := true
	array := []int{1, 2, 3}

	fmt.Println(getVariableType(num))      // Вывод: "int"
	fmt.Println(getVariableType(str))      // Вывод: "string"
	fmt.Println(getVariableType(floating)) // Вывод: "float64"
	fmt.Println(getVariableType(boolean))  // Вывод: "bool"
	fmt.Println(getVariableType(array))    // Вывод: "[]int"
}
