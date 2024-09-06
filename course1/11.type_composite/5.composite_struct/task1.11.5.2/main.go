package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

type Animal struct {
	Type string
	Name string
	Age  int
}

func getAnimals() []Animal {
	animals := []Animal{}
	for i := 0; i < 3; i++ {
		animals = append(animals, Animal{
			Type: gofakeit.Animal(),
			Name: gofakeit.Name(),
			Age:  gofakeit.Number(1, 15),
		})
	}
	return animals
}

func preparePrint(animals []Animal) string {
	result := ""
	for _, animal := range animals {
		result += fmt.Sprintf("Тип: %s, Имя: %s, Возраст: %d\n", animal.Type, animal.Name, animal.Age)
	}
	return result
}

func main() {
	gofakeit.Seed(0)

	animals := getAnimals()

	fmt.Print(preparePrint(animals))
}
