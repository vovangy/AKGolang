package main

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

func TestGenerateProducts(t *testing.T) {
	products := generateProducts(10)
	if len(products) != 10 {
		t.Errorf("Expected 10 products, got %d", len(products))
	}
	for _, p := range products {
		if p.Name == "" {
			t.Error("Product name should not be empty")
		}
		if p.Price < 1.0 || p.Price > 100.0 {
			t.Errorf("Product price %f is out of range", p.Price)
		}
		if p.CreatedAt.IsZero() {
			t.Error("Product CreatedAt should not be zero")
		}
		if p.Count < 1 || p.Count > 100 {
			t.Errorf("Product count %d is out of range", p.Count)
		}
	}
}

func TestSortByPrice(t *testing.T) {
	products := generateProducts(10)
	sort.Sort(ByPrice(products))

	for i := 1; i < len(products); i++ {
		if products[i-1].Price > products[i].Price {
			t.Errorf("Products not sorted by price. Index %d: %f > Index %d: %f", i-1, products[i-1].Price, i, products[i].Price)
		}
	}
}

func TestSortByCreatedAt(t *testing.T) {
	products := generateProducts(10)
	sort.Sort(ByCreatedAt(products))

	for i := 1; i < len(products); i++ {
		if products[i-1].CreatedAt.After(products[i].CreatedAt) {
			t.Errorf("Products not sorted by CreatedAt. Index %d: %v > Index %d: %v", i-1, products[i-1].CreatedAt, i, products[i].CreatedAt)
		}
	}
}

func TestSortByCount(t *testing.T) {
	products := generateProducts(10)
	sort.Sort(ByCount(products))

	for i := 1; i < len(products); i++ {
		if products[i-1].Count > products[i].Count {
			t.Errorf("Products not sorted by count. Index %d: %d > Index %d: %d", i-1, products[i-1].Count, i, products[i].Count)
		}
	}
}

func TestProductString(t *testing.T) {
	testDate := time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC)

	product := Product{
		Name:      "TestProduct",
		Price:     99.99,
		CreatedAt: testDate,
		Count:     10,
	}

	expected := fmt.Sprintf("Name: %s, Price: %.2f, Count: %d, CreatedAt: %s",
		product.Name, product.Price, product.Count, product.CreatedAt.String())

	if product.String() != expected {
		t.Errorf("Expected %q but got %q", expected, product.String())
	}
}
