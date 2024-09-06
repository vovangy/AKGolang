package main

import (
	"fmt"
	"strings"
)

func createUniqueText(text string) string {
	uniqueWords := map[string]struct{}{}
	arrayWords := []string{}

	words := strings.Fields(text)

	for _, word := range words {
		if _, ok := uniqueWords[word]; ok == false {
			uniqueWords[word] = struct{}{}
			arrayWords = append(arrayWords, word)
		}
	}

	separator := ""

	resultString := ""

	for _, value := range arrayWords {
		resultString += separator + value
		separator = " "
	}

	return resultString
}

func main() {
	fmt.Println(createUniqueText("bar bar bar foo foo baz"))
}
