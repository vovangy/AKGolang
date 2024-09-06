package main

import (
	"fmt"
	"os"
)

func getArgs() []string {
	if len(os.Args) > 1 {
		return os.Args[1:]
	}
	return nil
}

func main() {
	args := getArgs()
	fmt.Println("Position arguments:", args)
}
