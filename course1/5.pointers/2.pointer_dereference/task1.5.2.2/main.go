package main

import "fmt"

func Factorial(n *int) int {
	if n == nil || *n < 0 {
		return 0
	}

	res := 1
	for i := 1; i <= *n; i++ {
		res *= i
	}
	return res
}

func isPalindrome(str *string) bool {
	if str == nil {
		return false
	}

	start := 0
	end := len(*str) - 1
	result := true

	for end >= start {
		if (*str)[start] == (*str)[end] {
			start++
			end--
		} else {
			result = false
			break
		}
	}
	return result
}

func CountOccurrences(numbers *[]int, target *int) int {
	if numbers == nil || target == nil {
		return 0
	}

	count := 0
	for _, val := range *numbers {
		if val == *target {
			count++
		}
	}
	return count
}

func ReverseString(str *string) string {
	newString := ""
	for i := len(*str) - 1; i >= 0; i-- {
		newString += string((*str)[i])
	}
	return newString
}

func main() {
	var a int = 15
	fmt.Println(Factorial(&a))
	var b string = "123454321"
	fmt.Println(isPalindrome(&b))
	var c []int = []int{1, 2, 3, 4, 5, 3, 2}
	var target int = 3
	fmt.Println(CountOccurrences(&c, &target))
	var d string = "123456789"
	fmt.Println(ReverseString(&d))
}
