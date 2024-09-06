package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func WriteFile(data io.Reader, fd io.Writer) error {
	_, err := io.Copy(fd, data)
	if err != nil {
		return fmt.Errorf("не удалось записать данные в файл: %w", err)
	}
	return nil
}

func main() {
	filePath := "course1/13.popular_package/7.package_os/task1.13.7.2/file.txt"

	dir := filePath[:len(filePath)-len(findFileName(filePath))]
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Println("Ошибка при создании директории:", err)
		return
	}

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	err = WriteFile(strings.NewReader("Hello, World!"), file)
	if err != nil {
		fmt.Println("Ошибка при записи файла:", err)
		return
	}

	fmt.Println("Файл успешно записан.")
}

func findFileName(filePath string) string {
	for i := len(filePath) - 1; i >= 0; i-- {
		if filePath[i] == '/' {
			return filePath[i+1:]
		}
	}
	return filePath
}
