package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ctxWithoutCancel := context.WithoutCancel(ctx)

	go runBatch(ctx, "first")
	go runBatch(ctxWithoutCancel, "second")

	time.Sleep(5 * time.Second)

	fmt.Println("cancelling context")
	cancel()

	time.Sleep(time.Minute)
}

func runBatch(ctx context.Context, msg string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println(msg)
			time.Sleep(1 * time.Second)
		}
	}
}
