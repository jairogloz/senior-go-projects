package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	maxWorkers := 4

	tasks := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	taskChan := make(chan int, len(tasks))
	resultChan := make(chan bool, len(tasks))
	wg := &sync.WaitGroup{}

	for i := 0; i < maxWorkers; i++ {
		go worker(taskChan, resultChan, wg)
	}

	// Send tasks to workers
	for _, task := range tasks {
		wg.Add(1)
		taskChan <- task
	}
	close(taskChan)

	go func() {
		wg.Wait()
		// Cuando se cierra un canal, se pueden leer los datos ya enviados
		// pero no enviar nuevos datos
		close(resultChan)
	}()

	// Imprimir resultados
	for r := range resultChan {
		fmt.Println(r)
	}

}

func worker(taskChan <-chan int, resultChan chan<- bool, wg *sync.WaitGroup) {
	for task := range taskChan {
		time.Sleep(1 * time.Second)
		if task%3 == 0 {
			resultChan <- true
		} else {
			resultChan <- false
		}
		wg.Done()
	}
}
