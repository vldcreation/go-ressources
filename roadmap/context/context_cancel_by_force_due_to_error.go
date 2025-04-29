package main

import (
	"context"
	"fmt"
	"time"
)

func context_cancel_by_force_due_to_error(err error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(3*time.Second))

	chErr := make(chan error)

	go func(ctx context.Context) {
		// ... some process ...

		if err != nil {
			// cancel context by force, an error occurred
			chErr <- err
			return
		}

		// ... some other process ...

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
	case err := <-chErr:
		fmt.Println("process fail causing by some error:", err.Error())
	}
}
