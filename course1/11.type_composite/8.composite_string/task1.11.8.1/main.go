package main

import (
	"fmt"
	"unicode/utf8"
)

func countBytes(s string) int {
	return len(s)
}

func countSymbols(s string) int {
	return utf8.RuneCountInString(s)
}

func main() {
	bytes := countBytes("Привет, мир!")
	fmt.Println(bytes)

	symbols := countSymbols("Привет, мир!")
	fmt.Println(symbols)
}
