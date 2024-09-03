package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var builder strings.Builder
	builder.Grow(length)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charset))
		builder.WriteByte(charset[randomIndex])
	}

	return builder.String()
}

func main() {
	randomString := GenerateRandomString(10)
	fmt.Println(randomString)
}
