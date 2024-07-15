package main

import "fmt"

func CalculateStockValue(price float64, quantity int) (float64, float64) {
	totalValue := price * float64(quantity)
	return totalValue, price
}

func main() {
	fmt.Println(CalculateStockValue(300.50, 2))
}
