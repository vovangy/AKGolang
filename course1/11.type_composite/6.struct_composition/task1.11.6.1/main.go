package main

import "fmt"

type Dish struct {
	Name  string
	Price float64
}

type Order struct {
	Dishes []Dish
	Total  float64
}

func (order *Order) AddDish(dish Dish) {
	order.Dishes = append(order.Dishes, dish)
}

func (order *Order) RemoveDish(dish Dish) {
	if len(order.Dishes) <= 1 {
		order.Dishes = []Dish{}
		return
	}
	order.Dishes = order.Dishes[1:]
}

func (order *Order) CalculateTotal() {
	var Total float64 = 0
	for _, val := range order.Dishes {
		Total += val.Price
	}
	order.Total = Total
}

func main() {
	order := Order{}
	dish1 := Dish{Name: "Pizza", Price: 10.99}
	dish2 := Dish{Name: "Burger", Price: 5.99}

	order.AddDish(dish1)
	order.AddDish(dish2)

	order.CalculateTotal()
	fmt.Println("Total:", order.Total)

	order.RemoveDish(dish1)

	order.CalculateTotal()
	fmt.Println("Total:", order.Total)
}
