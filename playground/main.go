package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker function that processes tasks
func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		time.Sleep(time.Second) // Simulate work
		results <- task * 2     // Example processing (doubling the task value)
	}
}

func main() {
	numWorkers := 3
	numTasks := 10

	// Channels for task distribution and results
	taskChan := make(chan int, numTasks)
	resultChan := make(chan int, numTasks)
	var wg sync.WaitGroup

	// Start worker pool
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, taskChan, resultChan, &wg)
	}

	// Send tasks to workers
	for i := 1; i <= numTasks; i++ {
		taskChan <- i
	}
	close(taskChan) // Close task channel to signal workers that no more tasks will be sent

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect results
	for result := range resultChan {
		fmt.Println("Result:", result)
	}
}
