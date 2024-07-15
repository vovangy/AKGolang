package main

import (
	"fmt"
	"strings"
)

func UserInfo(name string, age int, cities ...string) string {
	resultString := "Имя: %s, возраст: %d, города: %s"
	citiesString := strings.Join(cities, ", ")
	return fmt.Sprintf(resultString, name, age, citiesString)
}

func main() {
	fmt.Println(UserInfo("John", 21, "Moscow", "Saint Petersburg"))
	fmt.Println(UserInfo("Alex", 34))
}
