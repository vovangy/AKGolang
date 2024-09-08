package main

import (
	"reflect"
	"testing"
	"time"
)

func TestNewOrder_Default(t *testing.T) {
	order := NewOrder(1)

	if order.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", order.ID)
	}

	if order.CustomerID != "" {
		t.Errorf("Expected CustomerID to be empty, got %s", order.CustomerID)
	}

	if len(order.Items) != 0 {
		t.Errorf("Expected Items to be empty, got %v", order.Items)
	}

	if !order.OrderDate.IsZero() {
		t.Errorf("Expected OrderDate to be zero, got %v", order.OrderDate)
	}
}

func TestNewOrder_WithCustomerID(t *testing.T) {
	order := NewOrder(1, WithCustomerID("123"))

	if order.CustomerID != "123" {
		t.Errorf("Expected CustomerID to be '123', got %s", order.CustomerID)
	}
}

func TestNewOrder_WithItems(t *testing.T) {
	items := []string{"item1", "item2"}
	order := NewOrder(1, WithItems(items))

	if !reflect.DeepEqual(order.Items, items) {
		t.Errorf("Expected Items to be %v, got %v", items, order.Items)
	}
}

func TestNewOrder_WithOrderDate(t *testing.T) {
	now := time.Now()
	order := NewOrder(1, WithOrderDate(now))

	if !order.OrderDate.Equal(now) {
		t.Errorf("Expected OrderDate to be %v, got %v", now, order.OrderDate)
	}
}

func TestNewOrder_AllOptions(t *testing.T) {
	now := time.Now()
	items := []string{"item1", "item2"}

	order := NewOrder(1,
		WithCustomerID("123"),
		WithItems(items),
		WithOrderDate(now),
	)

	if order.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", order.ID)
	}

	if order.CustomerID != "123" {
		t.Errorf("Expected CustomerID to be '123', got %s", order.CustomerID)
	}

	if !reflect.DeepEqual(order.Items, items) {
		t.Errorf("Expected Items to be %v, got %v", items, order.Items)
	}

	if !order.OrderDate.Equal(now) {
		t.Errorf("Expected OrderDate to be %v, got %v", now, order.OrderDate)
	}
}
