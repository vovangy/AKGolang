package main

import (
	"fmt"
	"strings"
)

func concatStrings(xs ...string) string {
	return strings.Join(xs, "")
}

func main() {
	result := concatStrings("Hello", " ", "world!")
	fmt.Println(result)

	fmt.Println(concatStrings())
	fmt.Println(concatStrings("Привет"))
	fmt.Println(concatStrings("Go", "lang"))
	fmt.Println(concatStrings("a", "", "b", "c"))
}
