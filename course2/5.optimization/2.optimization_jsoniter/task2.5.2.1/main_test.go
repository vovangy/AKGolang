package main

import (
	"testing"
)

func BenchmarkStandartJsonMarshal(b *testing.B) {
	users := GenerateUser(1000)
	serializer := &StandartJson{}
	for i := 0; i < b.N; i++ {
		for _, user := range users {
			_, _ = serializer.Marshal(&user)
		}
	}
}

func BenchmarkStandartJsonUnmarshal(b *testing.B) {
	users := GenerateUser(1000)
	serializer := &StandartJson{}
	data := make([][]byte, len(users))
	for i, user := range users {
		data[i], _ = serializer.Marshal(&user)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, d := range data {
			_, _ = serializer.Unmarshal(d)
		}
	}
}

func BenchmarkEasyJsonMarshal(b *testing.B) {
	users := GenerateUser(1000)
	serializer := &EasyJson{}
	for i := 0; i < b.N; i++ {
		for _, user := range users {
			_, _ = serializer.Marshal(&user)
		}
	}
}

func BenchmarkEasyJsonUnmarshal(b *testing.B) {
	users := GenerateUser(1000)
	serializer := &EasyJson{}
	data := make([][]byte, len(users))
	for i, user := range users {
		data[i], _ = serializer.Marshal(&user)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, d := range data {
			_, _ = serializer.Unmarshal(d)
		}
	}
}

// Бенчмарк для
func BenchmarkJsointerMarshal(b *testing.B) {
	users := GenerateUser(1000)
	serializer := &Jsointer{}
	for i := 0; i < b.N; i++ {
		for _, user := range users {
			_, _ = serializer.Marshal(&user)
		}
	}
}

func BenchmarkJsointerUnmarshal(b *testing.B) {
	users := GenerateUser(1000)
	serializer := &Jsointer{}
	data := make([][]byte, len(users))
	for i, user := range users {
		data[i], _ = serializer.Marshal(&user)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, d := range data {
			_, _ = serializer.Unmarshal(d)
		}
	}
}
