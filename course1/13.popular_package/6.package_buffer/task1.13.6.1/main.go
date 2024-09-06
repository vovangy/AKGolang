package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func getReader(b *bytes.Buffer) *bufio.Reader {
	return bufio.NewReader(b)
}

func main() {
	buffer := bytes.NewBufferString("Hello, World!")

	r := getReader(buffer)

	b := make([]byte, 13)
	_, err := r.Read(b)
	if err != nil {
		fmt.Println("Error reading from buffer:", err)
		return
	}

	fmt.Println(string(b))
}
