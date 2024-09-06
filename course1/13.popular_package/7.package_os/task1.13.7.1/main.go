package main

import (
	"fmt"
	"os"
)

func WriteFile(filePath string, data []byte, perm os.FileMode) error {
	dir := filePath[:len(filePath)-len(findFileName(filePath))]
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("не удалось создать директорию: %w", err)
	}

	err = os.WriteFile(filePath, data, perm)
	if err != nil {
		return fmt.Errorf("не удалось записать файл: %w", err)
	}

	return nil
}

func findFileName(filePath string) string {
	for i := len(filePath) - 1; i >= 0; i-- {
		if filePath[i] == '/' {
			return filePath[i+1:]
		}
	}
	return filePath
}

func main() {
	err := WriteFile("./path/to/file.txt", []byte("Hello, World!"), 0644)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Файл успешно записан.")
	}
}
