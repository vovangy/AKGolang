package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type Product struct {
	Name      string
	Price     float64
	CreatedAt time.Time
	Count     int
}

func (p Product) String() string {
	return fmt.Sprintf("Name: %s, Price: %.2f, Count: %d, CreatedAt: %s", p.Name, p.Price, p.Count, p.CreatedAt.String())
}

type ByPrice []Product

func (p ByPrice) Len() int           { return len(p) }
func (p ByPrice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByPrice) Less(i, j int) bool { return p[i].Price < p[j].Price }

type ByCreatedAt []Product

func (p ByCreatedAt) Len() int           { return len(p) }
func (p ByCreatedAt) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByCreatedAt) Less(i, j int) bool { return p[i].CreatedAt.Before(p[j].CreatedAt) }

type ByCount []Product

func (p ByCount) Len() int           { return len(p) }
func (p ByCount) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ByCount) Less(i, j int) bool { return p[i].Count < p[j].Count }

func generateProducts(n int) []Product {
	gofakeit.Seed(time.Now().UnixNano())
	products := make([]Product, n)
	for i := range products {
		products[i] = Product{
			Name:      gofakeit.Word(),
			Price:     gofakeit.Price(1.0, 100.0),
			CreatedAt: gofakeit.Date(),
			Count:     gofakeit.Number(1, 100),
		}
	}
	return products
}
