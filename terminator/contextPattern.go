package terminator

import (
	"context"
	"fmt"
	"time"
)

func infiniteLoop2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context cancelled")
			return
		default:
			fmt.Println("Running...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func RunWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	go infiniteLoop2(ctx)
	fmt.Println("Go Routine started")
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}
