package main

import (
	"errors"
	"fmt"
)

type Order interface {
	AddItem(item string, quantity int) error
	RemoveItem(item string) error
	GetOrderDetails() map[string]int
}

type DineInOrder struct {
	orderDetails map[string]int
}

func (d *DineInOrder) AddItem(item string, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}
	d.orderDetails[item] += quantity
	return nil
}

func (d *DineInOrder) RemoveItem(item string) error {
	var value int
	var exists bool

	if value, exists = d.orderDetails[item]; !exists {
		return errors.New("item does not exist in the order")
	}

	if value > 1 {
		d.orderDetails[item] = d.orderDetails[item] - 1
	} else {
		delete(d.orderDetails, item)
	}
	return nil
}

func (d *DineInOrder) GetOrderDetails() map[string]int {
	return d.orderDetails
}

type TakeAwayOrder struct {
	orderDetails map[string]int
}

func (t *TakeAwayOrder) AddItem(item string, quantity int) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than 0")
	}
	t.orderDetails[item] += quantity
	return nil
}

func (t *TakeAwayOrder) RemoveItem(item string) error {
	var value int
	var exists bool

	if value, exists = t.orderDetails[item]; !exists {
		return errors.New("item does not exist in the order")
	}

	if value > 1 {
		t.orderDetails[item] = t.orderDetails[item] - 1
	} else {
		delete(t.orderDetails, item)
	}
	return nil
}

func (t *TakeAwayOrder) GetOrderDetails() map[string]int {
	return t.orderDetails
}

func ManageOrder(o Order) {
	o.AddItem("Pizza", 2)
	o.AddItem("Burger", 1)
	o.RemoveItem("Pizza")
	o.RemoveItem("Pizza")
	fmt.Println(o.GetOrderDetails())
}

func main() {
	dineIn := &DineInOrder{orderDetails: make(map[string]int)}
	takeAway := &TakeAwayOrder{orderDetails: make(map[string]int)}

	fmt.Println("Dine-in Order:")
	ManageOrder(dineIn)

	fmt.Println("Takeaway Order:")
	ManageOrder(takeAway)
}
