package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Определение структур
type User struct {
	Name     string    `yaml:"name"`
	Age      int       `yaml:"age"`
	Comments []Comment `yaml:"comments"`
}

type Comment struct {
	Text string `yaml:"text"`
}

func writeYAML(filePath string, data []User) error {
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

	encoder := yaml.NewEncoder(file)
	defer encoder.Close()

	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("failed to encode YAML: %w", err)
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

	err := writeYAML("path/to/your/file.yaml", users)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
