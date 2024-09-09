package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
)

type Tabler interface {
	TableName() string
}

type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

func (u *User) TableName() string {
	return "users"
}

type SQLGenerator interface {
	CreateTableSQL(table Tabler) string
	CreateInsertSQL(model Tabler) string
}

type SQLiteGenerator struct{}

func (s *SQLiteGenerator) CreateTableSQL(table Tabler) string {
	t := reflect.TypeOf(table).Elem()
	var fields []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldName := field.Tag.Get("db_field")
		fieldType := field.Tag.Get("db_type")
		fields = append(fields, fmt.Sprintf("%s %s", fieldName, fieldType))
	}

	return fmt.Sprintf("CREATE TABLE %s (%s);", table.TableName(), strings.Join(fields, ", "))
}

func (s *SQLiteGenerator) CreateInsertSQL(model Tabler) string {
	t := reflect.TypeOf(model).Elem()
	v := reflect.ValueOf(model).Elem()
	var columns []string
	var values []string

	for i := 1; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldName := field.Tag.Get("db_field")
		columns = append(columns, fieldName)
		values = append(values, fmt.Sprintf("'%v'", v.Field(i).Interface()))
	}

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", model.TableName(), strings.Join(columns, ", "), strings.Join(values, ", "))
}

type FakeDataGenerator interface {
	GenerateFakeUser() User
}

type GoFakeitGenerator struct{}

func (g *GoFakeitGenerator) GenerateFakeUser() User {
	return User{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
	}
}

func main() {
	sqlGenerator := &SQLiteGenerator{}
	fakeDataGenerator := &GoFakeitGenerator{}
	user := User{}

	createTableSQL := sqlGenerator.CreateTableSQL(&user)
	fmt.Println(createTableSQL)

	for i := 0; i < 10; i++ {
		fakeUser := fakeDataGenerator.GenerateFakeUser()
		insertSQL := sqlGenerator.CreateInsertSQL(&fakeUser)
		fmt.Println(insertSQL)
	}
}
