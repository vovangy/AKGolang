package main

import (
	"fmt"
	"unicode/utf8"
)

func countUniqueUTF8Chars(s string) int {
	uniqueChars := make(map[string]struct{})

	var char []byte

	for i := 0; i < len(s); {
		_, size := utf8.DecodeRuneInString(s[i:])
		if size == 0 {
			i++
			continue
		}
		char = append(char[:0], s[i:i+size]...)
		uniqueChars[string(char)] = struct{}{}
		i += size
	}

	return len(uniqueChars)
}

func main() {
	fmt.Println(countUniqueUTF8Chars("Hello, World!")) // Вывод: 10
	fmt.Println(countUniqueUTF8Chars("Привет, мир!"))  // Вывод: 10
	fmt.Println(countUniqueUTF8Chars("Hello,    !"))   // Вывод: 9
}
