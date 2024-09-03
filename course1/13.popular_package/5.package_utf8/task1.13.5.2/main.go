package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func countRussianLetters(s string) map[rune]int {
	counts := make(map[rune]int)

	for i := 0; i < len(s); {
		char, size := utf8.DecodeRuneInString(s[i:])
		i += size

		if isRussianLetter(char) {
			char = unicode.ToLower(char)
			counts[char]++
		}
	}

	return counts
}

func isRussianLetter(char rune) bool {
	return (char >= 'А' && char <= 'Я') || (char >= 'а' && char <= 'я')
}

func main() {
	result := countRussianLetters("Привет, мир!")
	for key, value := range result {
		fmt.Printf("%c: %d\n", key, value)
	}
}
