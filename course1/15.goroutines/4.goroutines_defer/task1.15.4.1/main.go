package main

import (
	"fmt"
	"os"
)

func writeToFile(file *os.File, data string) error {
	defer file.Close()

	_, err := file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}

	err = writeToFile(file, "Hello, World!")
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
	} else {
		fmt.Println("Данные успешно записаны в файл")
	}
}
