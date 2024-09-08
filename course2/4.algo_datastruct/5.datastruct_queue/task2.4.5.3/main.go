package main

import (
	"container/list"
	"fmt"
)

type Car struct {
	LicensePlate string
}

type ParkingLot struct {
	space *list.List
}

func NewParkingLot() *ParkingLot {
	return &ParkingLot{space: list.New()}
}

func (p *ParkingLot) Park(c Car) {
	p.space.PushBack(c)
	fmt.Printf("Автомобиль %s припаркован.\n", c.LicensePlate)
}

func (p *ParkingLot) Leave() {
	if p.space.Len() == 0 {
		fmt.Println("Парковка пуста.")
		return
	}

	lastCar := p.space.Front()
	car := lastCar.Value.(Car)
	fmt.Printf("Автомобиль %s покинул парковку.\n", car.LicensePlate)

	p.space.Remove(lastCar)
}

func main() {
	parkingLot := NewParkingLot()

	parkingLot.Park(Car{LicensePlate: "ABC-123"})
	parkingLot.Park(Car{LicensePlate: "XYZ-789"})

	parkingLot.Leave()
	parkingLot.Leave()
	parkingLot.Leave()
}
