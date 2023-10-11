package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
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

	ctxHijo, cancelFunc := context.WithCancel(context.Background())

	go simulateLongOperation(ctxHijo)

	time.Sleep(3 * time.Second) // Give the operation time to complete or be canceled
	cancelFunc()

	(&mongo.Client{}).Database("ba").Collection("c").FindOne(ctxHijo, nil)

}
