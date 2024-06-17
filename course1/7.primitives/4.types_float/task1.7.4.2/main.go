package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mattevans/dinero"
)

const API_KEY = "fe6232c26cd54017b0e45ab9e829e96c"

func convertCurrency(fromCurrency, toCurrency string, amount float64) (float64, error) {
	client := dinero.NewClient(
		API_KEY,
		"USD",
		20*time.Minute,
	)

	rates, err := client.Rates.List()
	if err != nil {
		return 0, err
	}

	usdRate := rates.Rates["USD"]
	if fromCurrency != "USD" {
		fromRate, ok := rates.Rates[fromCurrency]
		if !ok {
			return 0, fmt.Errorf("rate for currency %s not found", fromCurrency)
		}
		amount = amount / fromRate
	}

	toRate, ok := rates.Rates[toCurrency]
	if !ok {
		return 0, fmt.Errorf("rate for currency %s not found", toCurrency)
	}

	convertedAmount := amount * (toRate / usdRate)
	return convertedAmount, nil
}

func main() {
	convertedAmount, err := convertCurrency("USD", "EUR", 100.0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(convertedAmount)
}
