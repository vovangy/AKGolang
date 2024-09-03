package main

import (
	"testing"
)

func BenchmarkWithoutPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		p = &Person{Age: i}
		_ = p
	}
}

func BenchmarkWithPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		p = NewPersonFromPool()
		p.Age = i
		PutPersonToPool(p)
	}
}
