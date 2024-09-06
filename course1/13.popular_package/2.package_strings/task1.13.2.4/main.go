package main

import (
	"math/rand"
	"strings"
	"time"
)

func generateActivationKey() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const partLength = 4
	const numParts = 4

	rand.Seed(time.Now().UnixNano())

	var parts []string

	for i := 0; i < numParts; i++ {
		var part strings.Builder
		for j := 0; j < partLength; j++ {
			randomIndex := rand.Intn(len(charset))
			part.WriteByte(charset[randomIndex])
		}
		parts = append(parts, part.String())
	}

	return strings.Join(parts, "-")
}
