package main

import "fmt"

func ConcatenateStrings(sep string, str ...string) string {
	resultOdd := "odd: "
	resultEven := "even: "
	curSep := ""
	for i := 0; i < len(str); i++ {
		if i > 1 {
			curSep = "-"
		}
		if i%2 == 0 {
			resultEven += curSep + str[i]
		} else {
			resultOdd += curSep + str[i]
		}
	}

	return resultEven + ", " + resultOdd
}

func main() {
	fmt.Println(ConcatenateStrings("-", "hello", "world", "how", "are", "you"))
}
