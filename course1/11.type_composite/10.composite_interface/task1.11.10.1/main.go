package main

import (
	"fmt"
)

func getType(i interface{}) string {
	if i == nil {
		return "Пустой интерфейс"
	}

	// return reflect.TypeOf(i).String()

	result := ""

	switch i.(type) {
	case int:
		result = "int"
	case string:
		result = "string"
	case []int:
		result = "[]int"
	default:
		result = "Неизвестный типы"
	}

	return result
}

func main() {
	var i interface{} = 42
	fmt.Println(getType(i))

	var j interface{} = "Hello, World!"
	fmt.Println(getType(j))

	var k interface{} = []int{1, 2, 3}
	fmt.Println(getType(k))

	var l interface{} = interface{}(nil)
	fmt.Println(getType(l))
}
