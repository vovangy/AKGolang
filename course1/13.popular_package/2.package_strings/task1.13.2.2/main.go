package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name     string
	Comments []Comment
}

type Comment struct {
	Message string
}

func main() {
	users := []User{
		{
			Name: "Betty",
			Comments: []Comment{
				{Message: "good Comment 1"},
				{Message: "BaD CoMmEnT 2"},
				{Message: "Bad Comment 3"},
				{Message: "Use camelCase please"},
			},
		},
		{
			Name: "Jhon",
			Comments: []Comment{
				{Message: "Good Comment 1"},
				{Message: "Good Comment 2"},
				{Message: "Good Comment 3"},
				{Message: "Bad Comments 4"},
			},
		},
	}

	users = FilterComments(users)
	fmt.Println(users)
}

func FilterComments(users []User) []User {
	var filteredUsers []User

	for _, user := range users {
		var filteredComments []Comment

		for _, comment := range user.Comments {
			if !IsBadComment(comment.Message) {
				filteredComments = append(filteredComments, comment)
			}
		}

		if len(filteredComments) > 0 {
			user.Comments = filteredComments
			filteredUsers = append(filteredUsers, user)
		}
	}

	return filteredUsers
}

func IsBadComment(comment string) bool {
	return strings.Contains(strings.ToLower(comment), "bad")
}

func GetBadComments(users []User) []Comment {
	var badComments []Comment

	for _, user := range users {
		for _, comment := range user.Comments {
			if IsBadComment(comment.Message) {
				badComments = append(badComments, comment)
			}
		}
	}

	return badComments
}

func GetGoodComments(users []User) []Comment {
	var goodComments []Comment

	for _, user := range users {
		for _, comment := range user.Comments {
			if !IsBadComment(comment.Message) {
				goodComments = append(goodComments, comment)
			}
		}
	}

	return goodComments
}
