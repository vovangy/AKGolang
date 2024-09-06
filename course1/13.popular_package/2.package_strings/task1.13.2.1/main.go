package main

import (
	"strings"
)

func CountWordsInText(txt string, words []string) map[string]int {
	wordCount := make(map[string]int)

	lowerTxt := strings.ToLower(txt)

	wordList := strings.Fields(lowerTxt)

	for i := range words {
		words[i] = strings.ToLower(words[i])
	}

	for _, word := range words {
		wordCount[word] = 0
	}

	for _, word := range wordList {
		if _, found := wordCount[word]; found {
			wordCount[word]++
		}
	}

	return wordCount
}
