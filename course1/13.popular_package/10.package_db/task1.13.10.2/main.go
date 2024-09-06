package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func PrepareQuery(operation, table string, user User) (string, []interface{}, error) {
	sb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question)
	var query string
	var args []interface{}
	var err error

	switch operation {
	case "create":
		query, args, err = sb.
			Insert(table).
			Columns("username", "email").
			Values(user.Username, user.Email).
			ToSql()
	case "insert":
		query, args, err = sb.
			Insert(table).
			Columns("username", "email").
			Values(user.Username, user.Email).
			ToSql()
	case "select":
		query, args, err = sb.
			Select("id", "username", "email").
			From(table).
			Where(squirrel.Eq{"id": user.ID}).
			ToSql()
	case "update":
		query, args, err = sb.
			Update(table).
			Set("username", user.Username).
			Set("email", user.Email).
			Where(squirrel.Eq{"id": user.ID}).
			ToSql()
	case "delete":
		query, args, err = sb.
			Delete(table).
			From(table).
			Where(squirrel.Eq{"id": user.ID}).
			ToSql()
	default:
		return "", nil, fmt.Errorf("invalid operation: %s", operation)
	}

	if err != nil {
		return "", nil, err
	}

	return query, args, nil
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
		username TEXT NOT NULL,
		email TEXT NOT NULL
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

	query, args, err := PrepareQuery("insert", "users", user)
	if err != nil {
		return err
	}

	_, err = db.Exec(query, args...)
	return err
}

func SelectUser(userID int) (User, error) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return User{}, err
	}
	defer db.Close()

	user := User{ID: userID}
	query, args, err := PrepareQuery("select", "users", user)
	if err != nil {
		return User{}, err
	}

	row := db.QueryRow(query, args...)
	err = row.Scan(&user.ID, &user.Username, &user.Email)
	if err == sql.ErrNoRows {
		return User{}, fmt.Errorf("user with id %d not found", userID)
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

	query, args, err := PrepareQuery("update", "users", user)
	if err != nil {
		return err
	}

	_, err = db.Exec(query, args...)
	return err
}

func DeleteUser(userID int) error {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return err
	}
	defer db.Close()

	user := User{ID: userID}
	query, args, err := PrepareQuery("delete", "users", user)
	if err != nil {
		return err
	}

	_, err = db.Exec(query, args...)
	return err
}

func main() {
	if err := CreateUserTable(); err != nil {
		log.Fatalf("Failed to create user table: %v", err)
	}

	newUser := User{Username: "alice", Email: "alice@example.com"}
	if err := InsertUser(newUser); err != nil {
		log.Fatalf("Failed to insert user: %v", err)
	}

	user, err := SelectUser(1)
	if err != nil {
		log.Fatalf("Failed to select user: %v", err)
	}
	fmt.Printf("Selected user: %+v\n", user)

	user.Email = "alice@newdomain.com"
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
