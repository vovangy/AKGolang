package main

import "fmt"

type User struct {
	Nickname string
	Age      int
	Email    string
}

func getUniqueUsers(users []User) []User {
	uniqueUsersNicknames := map[string]int{}

	for _, user := range users {
		uniqueUsersNicknames[user.Nickname] = uniqueUsersNicknames[user.Nickname] + 1
	}

	uniqueUsers := []User{}

	for _, user := range users {
		if uniqueUsersNicknames[user.Nickname] == 1 {
			uniqueUsers = append(uniqueUsers, user)
		}
	}

	return uniqueUsers
}

func main() {
	a := []User{User{Nickname: "aboba", Age: 5, Email: "vova"},
		User{Nickname: "aboba", Age: 5, Email: "vova"},
		User{Nickname: "aboba", Age: 5, Email: "vova"},
		User{Nickname: "aboba", Age: 5, Email: "vova"},
		User{Nickname: "zeliboba", Age: 5, Email: "vova"}}
	b := getUniqueUsers(a)

	fmt.Println(b)
}
