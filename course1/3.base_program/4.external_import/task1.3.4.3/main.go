package main

import (
	"fmt"

	"github.com/icrowley/fake"
)

func GenerateFakeData() string {
	fullname := fake.FullName()
	email := fake.EmailAddress()
	phone := fake.Phone()
	address := fake.StreetAddress()
	return fmt.Sprintf("Name: %s\nAddress: %s\nPhone: %s\nEmail: %s", fullname, address, phone, email)

}

func main() {
	fmt.Println(GenerateFakeData())
}
