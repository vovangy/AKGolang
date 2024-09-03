package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Cache struct {
	data sync.Map
}

func (c *Cache) Set(key string, value interface{}) {
	c.data.Store(key, value)
}

func (c *Cache) Get(key string) (interface{}, bool) {
	return c.data.Load(key)
}

func main() {
	cache := &Cache{}

	for i := 0; i < 100; i++ {
		cache.Set(strconv.Itoa(i), fmt.Sprintf("user-%d", i))
	}

	for i := 0; i < 100; i++ {
		go func(i int) {
			if value, ok := cache.Get(strconv.Itoa(i)); ok {
				fmt.Println(value)
			} else {
				fmt.Printf("Key %d not found\n", i)
			}
		}(i)
	}

	time.Sleep(1 * time.Second)
}
