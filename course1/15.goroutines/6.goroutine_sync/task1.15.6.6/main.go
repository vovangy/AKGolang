package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type User struct {
	ID   int
	Name string
}

type Cache struct {
	mu    sync.RWMutex
	store map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{
		store: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = value
}

func (c *Cache) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.store[key]
}

func keyBuilder(keys ...string) string {
	return strings.Join(keys, ":")
}

func GetUser(i interface{}) *User {
	return i.(*User)
}

func main() {
	cache := NewCache()

	for i := 0; i < 100; i++ {
		go cache.Set(keyBuilder("user", strconv.Itoa(i)), &User{
			ID:   i,
			Name: fmt.Sprintf("user-%d", i),
		})
	}

	for i := 0; i < 100; i++ {
		go func(i int) {
			raw := cache.Get(keyBuilder("user", strconv.Itoa(i)))
			fmt.Println(GetUser(raw))
		}(i)
	}

	time.Sleep(1 * time.Second)
}
