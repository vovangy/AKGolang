package main

import "fmt"

func ReplaceSymbols(str string, old, new rune) string {
	newString := ""

	for _, sym := range str {
		if sym == old {
			newString += string(new)
		} else {
			newString += string(sym)
		}
	}

	return newString
}

func main() {
	result := ReplaceSymbols("Hello, world!", 'o', '0')
	fmt.Println(result)
}
