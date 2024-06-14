package main

import (
	"fmt"
	"math"
)

func Add(a, b int) *int {
	res := a + b
	return &res
}

func Max(numbers []int) *int {
	if len(numbers) == 0 {
		return nil
	}

	max := numbers[0]

	for _, val := range numbers {
		if max < val {
			max = val
		}
	}

	return &max
}

func IsPrime(number int) *bool {
	high := int(math.Sqrt(float64(number)))
	isPrime := true
	for i := 2; i <= high; i++ {
		if number%i == 0 {
			isPrime = false
			break
		}
	}

	if number == 1 {
		isPrime = false
	}

	return &isPrime
}

func ConcatenateStrings(strs []string) *string {
	if strs == nil {
		return nil
	}

	resultString := ""

	for _, val := range strs {
		resultString += val
	}

	return &resultString
}

func main() {
	fmt.Println(*Add(1, 9))
	fmt.Println(*Max([]int{1, 2, 9, 3, 4, 5}))
	fmt.Println(*IsPrime(96))
	fmt.Println(*ConcatenateStrings([]string{"Kto ", "Zdes ", "? ", ":O"}))
	return
}
