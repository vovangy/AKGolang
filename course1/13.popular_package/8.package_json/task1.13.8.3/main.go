package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Comment struct {
	Text string `json:"text"`
}

type User struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Comments []Comment `json:"comments"`
}

func writeJSON(filePath string, data []User) error {
	dir := filepath.Dir(filePath)

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory '%s': %w", dir, err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("failed to write data to file: %w", err)
	}

	return nil
}

func main() {
	users := []User{
		{
			Name: "John",
			Age:  30,
			Comments: []Comment{
				{Text: "Great post!"},
				{Text: "I agree"},
			},
		},
		{
			Name: "Alice",
			Age:  25,
			Comments: []Comment{
				{Text: "Nice article"},
			},
		},
	}

	filePath := "data/users.json"

	err := writeJSON(filePath, users)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Data successfully written to", filePath)
}
