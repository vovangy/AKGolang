package main

import (
	"fmt"
	"strings"
)

func filterSentence(sentence string, filter map[string]bool) string {
	words := strings.Fields(sentence)
	resultString := ""
	separator := ""

	for _, val := range words {
		if _, ok := filter[val]; !ok {
			resultString += separator + val
			separator = " "
		}
	}

	return resultString
}

func main() {
	sentence := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	filter := map[string]bool{"ipsum": true, "elit": true}

	filteredSentence := filterSentence(sentence, filter)

	fmt.Println(filteredSentence)
}
