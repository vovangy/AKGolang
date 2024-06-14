package main

import "fmt"

func mutate(a *int) {
	*a = 42
}

func ReverseString(a *string) {
	newString := ""
	for i := len(*a) - 1; i >= 0; i-- {
		newString += string((*a)[i])
	}
	*a = newString
}

func main() {
	var a int = 11
	mutate(&a)
	fmt.Println(a)
	var str string = "abcde"
	ReverseString(&str)
	fmt.Println(str)
}
