package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Number of cores: ", runtime.NumCPU())
	maxWorkers := 9
	tasks := []func() int{
		func() int { return 1 },
		func() int { return 2 },
		func() int { return 3 },
		func() int { panic("error") }, // Example of a task that panics
		func() int { return 4 },
		func() int { return 5 },
		func() int { panic("error") }, // Example of a task that panics
		func() int { return 6 },
		func() int { return 7 },
		func() int { return 8 },
		func() int { return 9 },
		func() int { return 10 },
		func() int { return 11 },
		func() int { return 12 },
		func() int { return 13 },
		func() int { return 14 },
		func() int { return 15 },
	}

	taskChan := make(chan func() int, len(tasks))
	resultChan := make(chan int, len(tasks))
	wg := &sync.WaitGroup{}

	// spawn maxWorkers
	for i := 0; i < maxWorkers; i++ {
		go worker(taskChan, resultChan, wg)
	}

	for _, task := range tasks {
		wg.Add(1)
		taskChan <- task
	}
	close(taskChan)

	go func() {
		wg.Wait()
		close(resultChan)
		fmt.Println("closed result chan")
	}()

	output := make([]int, 0)
	for r := range resultChan {
		output = append(output, r)
	}

	if len(output) == 0 {
		fmt.Println("No results")
	} else {
		fmt.Println("Results gathered: ", output)
	}
}

func worker(taskChan <-chan func() int, resultChan chan<- int, wg *sync.WaitGroup) {
	for task := range taskChan {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered from panic: ", r)
				}
				wg.Done()
			}()
			resultChan <- task()
		}()
	}
}
