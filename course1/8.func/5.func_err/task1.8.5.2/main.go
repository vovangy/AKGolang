package main

import "fmt"

func CheckDiscount(price, discount float64) (float64, error) {
	if discount > 50 {
		return 0, fmt.Errorf("Скидка не может превышать 50%")
	}

	return (100 - discount) / 100 * price, nil
}

func main() {
	fmt.Println(CheckDiscount(9000, 60))
}
