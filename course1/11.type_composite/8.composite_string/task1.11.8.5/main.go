package main

import (
	"fmt"
)

func ReverseString(str string) string {
	runes := []rune(str)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	fmt.Println(ReverseString("Hello, world!"))
	fmt.Println(ReverseString("12345"))
	fmt.Println(ReverseString(""))
	fmt.Println(ReverseString("a"))
	fmt.Println(ReverseString("   "))
	fmt.Println(ReverseString("!@#$%^&*()_+"))
}
