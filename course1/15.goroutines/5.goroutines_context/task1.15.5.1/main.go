package main

import (
	"context"
	"fmt"
	"time"
)

func contextWithDeadline(ctx context.Context, contextDeadline time.Duration, timeAfter time.Duration) string {
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(contextDeadline))
	defer cancel()

	select {
	case <-ctx.Done():
		return "context deadline exceeded"
	case <-time.After(timeAfter):
		return "time after exceeded"
	}
}

func main() {
	res := contextWithDeadline(context.Background(), 2*time.Second, 1*time.Second)
	fmt.Println(res)

	res = contextWithDeadline(context.Background(), 1*time.Second, 2*time.Second)
	fmt.Println(res)
}
