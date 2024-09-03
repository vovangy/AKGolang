package main

import (
	"fmt"
	"strings"
)

func generateMathString(operands []int, operator string) string {
	var result int
	var operationStrings []string

	if len(operands) > 0 {
		result = operands[0]
		operationStrings = append(operationStrings, fmt.Sprintf("%d", operands[0]))
	}

	for _, operand := range operands[1:] {
		operationStrings = append(operationStrings, fmt.Sprintf("%s %d", operator, operand))
		switch operator {
		case "+":
			result += operand
		case "-":
			result -= operand
		case "*":
			result *= operand
		case "/":
			if operand != 0 {
				result /= operand
			}
		}
	}

	finalString := fmt.Sprintf("%s = %d", strings.Join(operationStrings, " "), result)
	return finalString
}

func main() {
	fmt.Println(generateMathString([]int{2, 4, 6}, "+"))  // "2 + 4 + 6 = 12"
	fmt.Println(generateMathString([]int{10, 5, 2}, "-")) // "10 - 5 - 2 = 3"
	fmt.Println(generateMathString([]int{3, 4, 5}, "*"))  // "3 * 4 * 5 = 60"
	fmt.Println(generateMathString([]int{20, 4, 2}, "/")) // "20 / 4 / 2 = 2"
}
