package main

import "fmt"

func main() {
	var name string
	var age int
	var town string

	fmt.Print("Введите ваше имя: ")
	fmt.Scanln(&name)
	fmt.Print("Введите ваш возраст: ")
	fmt.Scanln(&age)
	fmt.Print("Введите ваш город: ")
	fmt.Scanln(&town)

	fmt.Printf("Имя: %s\nВозраст: %d\nГород: %s\n", name, age, town)
}
