package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func CreateUserTable() error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER NOT NULL
	);
	`
	_, err = db.Exec(query)
	return err
}

func InsertUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query := `INSERT INTO users (name, age) VALUES (?, ?)`
	_, err = db.Exec(query, user.Name, user.Age)
	return err
}

func SelectUser(id int) (User, error) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	query := `SELECT id, name, age FROM users WHERE id = ?`
	row := db.QueryRow(query, id)

	var user User
	err = row.Scan(&user.ID, &user.Name, &user.Age)
	if err == sql.ErrNoRows {
		return User{}, fmt.Errorf("user with id %d not found", id)
	} else if err != nil {
		return User{}, err
	}

	return user, nil
}

func UpdateUser(user User) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query := `UPDATE users SET name = ?, age = ? WHERE id = ?`
	_, err = db.Exec(query, user.Name, user.Age, user.ID)
	return err
}

func DeleteUser(id int) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	query := `DELETE FROM users WHERE id = ?`
	_, err = db.Exec(query, id)
	return err
}

func main() {
	if err := CreateUserTable(); err != nil {
		log.Fatalf("Failed to create user table: %v", err)
	}

	newUser := User{Name: "Alice", Age: 30}
	if err := InsertUser(newUser); err != nil {
		log.Fatalf("Failed to insert user: %v", err)
	}

	user, err := SelectUser(1)
	if err != nil {
		log.Fatalf("Failed to select user: %v", err)
	}
	fmt.Printf("Selected user: %+v\n", user)

	user.Age = 31
	if err := UpdateUser(user); err != nil {
		log.Fatalf("Failed to update user: %v", err)
	}

	user, err = SelectUser(1)
	if err != nil {
		log.Fatalf("Failed to select user: %v", err)
	}
	fmt.Printf("Selected user: %+v\n", user)

	if err := DeleteUser(1); err != nil {
		log.Fatalf("Failed to delete user: %v", err)
	}
}
