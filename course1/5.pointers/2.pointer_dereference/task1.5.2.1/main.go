package main

import "fmt"

func Dereference(n *int) int {
	return *n
}

func Sum(a, b *int) int {
	return *a + *b
}

func main() {
	var a int = 99
	fmt.Println(Dereference(&a))
	var b int = 101
	fmt.Println(Sum(&a, &b))
}
