package main

import (
	"context"
	"fmt"
	"time"
)

func contextWithTimeout(ctx context.Context, contextTimeout time.Duration, timeAfter time.Duration) string {
	ctx, cancel := context.WithTimeout(ctx, contextTimeout)
	defer cancel()

	select {
	case <-ctx.Done():
		return "context timeout exceeded"
	case <-time.After(timeAfter):
		return "time after exceeded"
	}
}

func main() {
	res := contextWithTimeout(context.Background(), 2*time.Second, 1*time.Second)
	fmt.Println(res)

	res = contextWithTimeout(context.Background(), 1*time.Second, 2*time.Second)
	fmt.Println(res)
}
