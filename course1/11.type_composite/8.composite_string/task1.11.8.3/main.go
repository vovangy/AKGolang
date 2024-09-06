package main

import (
	"fmt"
)

func getBytes(s string) []byte {
	return []byte(s)
}

func getRunes(s string) []rune {
	return []rune(s)
}

func main() {
	s := "Привет, мир!"

	bytes := getBytes(s)
	fmt.Printf("Bytes: %v\n", bytes)

	runes := getRunes(s)
	fmt.Printf("Runes: %v\n", runes)
}
