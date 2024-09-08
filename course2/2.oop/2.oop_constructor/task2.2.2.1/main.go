package main

import (
	"time"
)

type Order struct {
	ID         int
	CustomerID string
	Items      []string
	OrderDate  time.Time
}

type OrderOption func(*Order)

func NewOrder(id int, opts ...OrderOption) *Order {
	order := &Order{
		ID: id,
	}
	for _, opt := range opts {
		opt(order)
	}
	return order
}

func WithCustomerID(customerID string) OrderOption {
	return func(o *Order) {
		o.CustomerID = customerID
	}
}

func WithItems(items []string) OrderOption {
	return func(o *Order) {
		o.Items = items
	}
}

func WithOrderDate(orderDate time.Time) OrderOption {
	return func(o *Order) {
		o.OrderDate = orderDate
	}
}
