package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

type Person struct {
	Name string `json:"name" yaml:"name"`
	Age  int    `json:"age" yaml:"age"`
}

func unmarshal(data []byte, v interface{}) error {
	if err := json.Unmarshal(data, v); err == nil {
		return nil
	}

	if err := yaml.Unmarshal(data, v); err == nil {
		return nil
	}

	return fmt.Errorf("failed to unmarshal data as JSON or YAML")
}

func main() {
	jsonData := []byte(`{"name": "John", "age": 30}`)
	var personFromJSON Person
	if err := unmarshal(jsonData, &personFromJSON); err != nil {
		fmt.Println("Ошибка декодирования JSON данных:", err)
		return
	}
	fmt.Println("Из JSON - Имя:", personFromJSON.Name)
	fmt.Println("Из JSON - Возраст:", personFromJSON.Age)

	yamlData := []byte(`
name: Alice
age: 25
`)
	var personFromYAML Person
	if err := unmarshal(yamlData, &personFromYAML); err != nil {
		fmt.Println("Ошибка декодирования YAML данных:", err)
		return
	}
	fmt.Println("Из YAML - Имя:", personFromYAML.Name)
	fmt.Println("Из YAML - Возраст:", personFromYAML.Age)
}
