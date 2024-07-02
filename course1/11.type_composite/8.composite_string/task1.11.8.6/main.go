package main

import (
	"fmt"
)

func CountVowels(str string) int {
	vowels := "аеёиоуыэюяАЕЁИОУЫЭЮЯaeiouAEIOU"
	vowelSet := make(map[rune]struct{}, len(vowels))
	for _, v := range vowels {
		vowelSet[v] = struct{}{}
	}

	count := 0
	for _, char := range str {
		if _, exists := vowelSet[char]; exists {
			count++
		}
	}

	return count
}

func main() {
	count := CountVowels("Привет, мир!")
	fmt.Println(count)

	count = CountVowels("Hello, world!")
	fmt.Println(count)

	count = CountVowels("bcdfg")
	fmt.Println(count)

	count = CountVowels("AEIOUaeiou")
	fmt.Println(count)
}
