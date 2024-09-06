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
	store map[string]*User
}

func NewCache() *Cache {
	return &Cache{
		store: make(map[string]*User),
	}
}

func (c *Cache) Set(key string, user *User) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = user
}

func (c *Cache) Get(key string) *User {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.store[key]
}

func keyBuilder(keys ...string) string {
	return strings.Join(keys, ":")
}

func main() {
	cache := NewCache()

	for i := 0; i < 100; i++ {
		go cache.Set(keyBuilder("user", strconv.Itoa(i)), &User{
			ID:   i,
			Name: fmt.Sprintf("user-%d", i),
		})
	}

	time.Sleep(1 * time.Second)

	for i := 0; i < 100; i++ {
		go func(i int) {
			user := cache.Get(keyBuilder("user", strconv.Itoa(i)))
			if user != nil {
				fmt.Println(user)
			} else {
				fmt.Printf("user-%d not found\n", i)
			}
		}(i)
	}

	time.Sleep(1 * time.Second)
}
