package main

import (
	"fmt"
	"os"
)

// ReadString читает содержимое файла и возвращает его в виде строки
func ReadString(filePath string) (string, error) {
	// Открываем файл для чтения
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("не удалось открыть файл: %w", err)
	}
	defer file.Close() // Закрываем файл после завершения работы функции

	// Читаем содержимое файла
	stat, err := file.Stat()
	if err != nil {
		return "", fmt.Errorf("не удалось получить информацию о файле: %w", err)
	}

	// Создаем буфер для хранения данных
	data := make([]byte, stat.Size())

	// Читаем данные из файла
	_, err = file.Read(data)
	if err != nil {
		return "", fmt.Errorf("не удалось прочитать файл: %w", err)
	}

	// Возвращаем содержимое файла как строку
	return string(data), nil
}

func main() {
	filePath := "example.txt"

	content, err := ReadString(filePath)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	fmt.Println("Содержимое файла:")
	fmt.Println(content)
}
