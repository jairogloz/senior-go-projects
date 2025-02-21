package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobsChan <-chan int, resultChan chan<- int, wg *sync.WaitGroup) {
	for job := range jobsChan {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered from panic: ", r)
				}
				wg.Done()
			}()

			time.Sleep(1 * time.Second)
			fmt.Printf("Worker %d processing job %d\n", id, job)
			if job == 1 {
				panic("hell no!")
			}
			result := job * job
			resultChan <- result
		}()
	}
}

func main() {

	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	jobsChan := make(chan int, 10)
	resultChan := make(chan int, 10)

	wg := &sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		go worker(i, jobsChan, resultChan, wg)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for _, job := range jobs {
		wg.Add(1)
		jobsChan <- job
	}
	close(jobsChan)

	for result := range resultChan {
		fmt.Println("result: ", result)
	}

}
