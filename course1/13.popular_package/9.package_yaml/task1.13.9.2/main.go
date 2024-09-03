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

func getConfigFromYAML(data []byte) (Config, error) {
	var config Config

	err := yaml.Unmarshal(data, &config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal YAML data: %w", err)
	}

	return config, nil
}

func main() {
	yamlData := `
server:
  port: "8080"
db:
  host: "localhost"
  port: "5432"
  user: "admin"
  password: "password123"
`

	config, err := getConfigFromYAML([]byte(yamlData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Server Port: %s\n", config.Server.Port)
	fmt.Printf("DB Host: %s\n", config.Db.Host)
	fmt.Printf("DB Port: %s\n", config.Db.Port)
	fmt.Printf("DB User: %s\n", config.Db.User)
	fmt.Printf("DB Password: %s\n", config.Db.Password)
}
