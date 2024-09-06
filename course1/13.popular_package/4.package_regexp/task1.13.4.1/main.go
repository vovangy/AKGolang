package main

import (
	"fmt"
	"regexp"
)

func isValidEmail(email string) bool {
	const emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailPattern)
	return re.MatchString(email)
}

func main() {
	emails := []string{
		"test@example.com",
		"invalid_email",
		"user@sub.domain.com",
		"another.test@domain.co.uk",
		"invalid-email@",
		"@missingusername.com",
	}

	for _, email := range emails {
		valid := isValidEmail(email)
		if valid {
			fmt.Printf("%s является валидным email-адресом\n", email)
		} else {
			fmt.Printf("%s не является валидным email-адресом\n", email)
		}
	}
}
