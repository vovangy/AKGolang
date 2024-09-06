package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func getScanner(b *bytes.Buffer) *bufio.Scanner {
	return bufio.NewScanner(b)
}

func main() {
	data := []byte("Hello\n,\n World!")
	buffer := bytes.NewBuffer(data)

	scanner := getScanner(buffer)

	if scanner == nil {
		panic("Expected non-nil scanner, got nil")
	}

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error during scanning:", err)
	}
}
