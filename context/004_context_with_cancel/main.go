package main

import (
	"context"
	"fmt"
	"time"
)

func waitForEventOrCancel(ctx context.Context) {
	// Create a select statement to wait on multiple operations.
	select {
	case <-time.After(5 * time.Second):
		// If 5 seconds pass, print that the operation completed normally.
		fmt.Println("Completed without cancellation")
	case <-ctx.Done():
		// If the context is cancelled, print the cancellation message.
		fmt.Println("Cancelled:", ctx.Err())
	}
}

func main() {
	// Create a context with cancellation functionality.
	ctx, cancel := context.WithCancel(context.Background())

	// Start the goroutine, passing the context.
	go waitForEventOrCancel(ctx)

	// Wait for 3 seconds before cancelling the context.
	time.Sleep(3 * time.Second)
	cancel() // This will trigger the ctx.Done() case in the goroutine.

	// Wait a little bit to see the output after cancellation.
	time.Sleep(1 * time.Second)
}
