package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func writeJSON(filePath string, data interface{}) error {
	dir := filepath.Dir(filePath)

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory '%s': %w", dir, err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file '%s': %w", filePath, err)
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("failed to write data to file '%s': %w", filePath, err)
	}

	return nil
}

func main() {
	// Пример данных
	data := []map[string]interface{}{
		{
			"name": "Elliot",
			"age":  25,
		},
		{
			"name": "Fraser",
			"age":  30,
		},
	}

	filePath := "users.json"

	err := writeJSON(filePath, data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Data successfully written to", filePath)
}
