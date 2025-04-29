package main

import (
	"context"
	"fmt"
	"time"
)

func context_cancel_by_force_cancle_func() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(3*time.Second))

	go func(ctx context.Context) {
		// simulate a process that takes 2 seconds to complete
		time.Sleep(2 * time.Second)

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
