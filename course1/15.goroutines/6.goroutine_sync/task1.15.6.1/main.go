package main

import (
	"sync"
)

type Person struct {
	Age int
}

var personPool = sync.Pool{
	New: func() interface{} {
		return &Person{}
	},
}

func NewPersonFromPool() *Person {
	return personPool.Get().(*Person)
}

func PutPersonToPool(p *Person) {
	personPool.Put(p)
}
