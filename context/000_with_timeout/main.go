package main

import (
	"context"
	"fmt"
	"time"
)

// simulateLongOperation simulates a long-running operation by sleeping for 2 seconds.
func simulateLongOperation(ctx context.Context) {
	for i := 0; i < 5; i++ {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Doing some work...")
		case <-ctx.Done():
			fmt.Println("Operation canceled")
			return
		}
	}
	fmt.Println("Operation completed successfully")
}

func main() {

	ctxHijo, _ := context.WithTimeout(context.Background(), 3*time.Second)

	go simulateLongOperation(ctxHijo)

	time.Sleep(6 * time.Second) // Give the operation time to complete or be canceled

}
