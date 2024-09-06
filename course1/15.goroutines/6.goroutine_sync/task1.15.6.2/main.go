package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func waitGroupExample(goroutines ...func(n int) string) string {
	var wg sync.WaitGroup
	var mu sync.Mutex
	results := make([]string, len(goroutines))

	wg.Add(len(goroutines))

	for i, gr := range goroutines {
		go func(i int, gr func(n int) string) {
			defer wg.Done()
			result := gr(i)
			time.Sleep(time.Second * time.Duration(rand.Intn(3)))
			mu.Lock()
			results[i] = result
			mu.Unlock()
		}(i, gr)
	}

	wg.Wait()
	return strings.Join(results, "\n")
}

func main() {
	count := 1000
	goroutines := make([]func(n int) string, count)

	for i := 0; i < count; i++ {
		j := i
		goroutines[i] = func(n int) string {
			return fmt.Sprintf("goroutine %d done", j)
		}
	}

	fmt.Println(waitGroupExample(goroutines...))
}
