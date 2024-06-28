package main

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

type User struct {
	Name string
	Age  int
}

func getUsers() []User {
	users := make([]User, 10)
	for i := 0; i < 10; i++ {
		users[i] = User{
			Name: gofakeit.Name(),
			Age:  gofakeit.Number(18, 60),
		}
	}
	return users
}

func preparePrint(users []User) string {
	result := ""
	for _, user := range users {
		result += fmt.Sprintf("Имя: %s, Возраст: %d\n", user.Name, user.Age)
	}
	return result
}

func main() {
	gofakeit.Seed(0)

	users := getUsers()

	resString := preparePrint(users)

	fmt.Print(resString)

}
