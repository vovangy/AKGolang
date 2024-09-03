package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server Server `yaml:"server"`
	Db     Db     `yaml:"db"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Db struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func getYAML(configs []Config) (string, error) {
	yamlData, err := yaml.Marshal(configs)
	if err != nil {
		return "", fmt.Errorf("failed to marshal data to YAML: %w", err)
	}

	return string(yamlData), nil
}

func main() {
	configs := []Config{
		{
			Server: Server{
				Port: "8080",
			},
			Db: Db{
				Host:     "localhost",
				Port:     "5432",
				User:     "admin",
				Password: "password123",
			},
		},
	}

	yamlStr, err := getYAML(configs)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(yamlStr)
}
