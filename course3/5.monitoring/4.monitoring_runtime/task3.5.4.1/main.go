package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"time"

	"golang.org/x/sync/errgroup"
)

func monitorGoroutines(prevGoroutines int) {
	currentGoroutines := prevGoroutines

	for {
		time.Sleep(300 * time.Millisecond)
		goroutinesNow := runtime.NumGoroutine()
		goroutinesDifference := float64(goroutinesNow) / float64(currentGoroutines)

		if goroutinesDifference < 0.8 {
			fmt.Println("Number of goroutines discrese for more than 20%%")
		} else if goroutinesDifference > 1.2 {
			fmt.Println("Number of goroutines increse for more that 20%%")
		}

		fmt.Printf("Now it is %d goroutines\n", goroutinesNow)
		currentGoroutines = goroutinesNow
	}
}
func main() {
	g, _ := errgroup.WithContext(context.Background())

	go func() {
		monitorGoroutines(runtime.NumGoroutine())
	}()

	for i := 0; i < 64; i++ {
		g.Go(func() error {
			time.Sleep(5 * time.Second)
			return nil
		})
		time.Sleep(80 * time.Millisecond)
	}

	if err := g.Wait(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
