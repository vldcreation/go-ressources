package main

import (
	"context"
	"fmt"
	"time"
)

func context_timeout_exceeded() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(3*time.Second))

	go func(ctx context.Context) {
		// simulate a process that takes 4 second to complete
		time.Sleep(4 * time.Second)

		// cancel context by force, assuming the whole process is complete
		cancel()
	}(ctx)

	select {
	case <-ctx.Done():
		switch ctx.Err() {
		case context.DeadlineExceeded:
			fmt.Println("context timeout exceeded")
		case context.Canceled:
			fmt.Println("context cancelled by force. whole process is complete")
		}
	}
}
