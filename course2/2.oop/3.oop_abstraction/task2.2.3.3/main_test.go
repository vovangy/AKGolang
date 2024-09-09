package main

import "testing"

func TestUser_TableName(t *testing.T) {
	user := &User{}
	if user.TableName() != "users" {
		t.Errorf("Expected table name 'users', but got %s", user.TableName())
	}
}

func TestSQLiteGenerator_CreateTableSQL(t *testing.T) {
	sqlGen := &SQLiteGenerator{}
	user := &User{}

	expectedSQL := "CREATE TABLE users (id SERIAL PRIMARY KEY, first_name VARCHAR(100), last_name VARCHAR(100), email VARCHAR(100) UNIQUE);"
	actualSQL := sqlGen.CreateTableSQL(user)

	if actualSQL != expectedSQL {
		t.Errorf("Expected SQL: %s, but got %s", expectedSQL, actualSQL)
	}
}

func TestSQLiteGenerator_CreateInsertSQL(t *testing.T) {
	sqlGen := &SQLiteGenerator{}
	user := &User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}

	expectedSQL := "INSERT INTO users (first_name, last_name, email) VALUES ('John', 'Doe', 'john.doe@example.com');"
	actualSQL := sqlGen.CreateInsertSQL(user)

	if actualSQL != expectedSQL {
		t.Errorf("Expected SQL: %s, but got %s", expectedSQL, actualSQL)
	}
}

func TestGoFakeitGenerator_GenerateFakeUser(t *testing.T) {
	fakeDataGen := &GoFakeitGenerator{}
	user := fakeDataGen.GenerateFakeUser()

	if user.FirstName == "" || user.LastName == "" || user.Email == "" {
		t.Error("Generated user contains empty fields")
	}
}
