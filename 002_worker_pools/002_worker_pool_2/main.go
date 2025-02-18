package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type txResult struct {
	txID    int
	success bool
	err     error
}

func main() {

	transactions := []int{101, 102, 103, 104, 105, 106, 107, 108, 109}
	maxWorkers := 3

	txChan := make(chan int, len(transactions))
	resultChan := make(chan txResult, len(transactions))

	wg := &sync.WaitGroup{}

	for i := 0; i < maxWorkers; i++ {
		go worker(txChan, resultChan, wg)
	}

	for _, txID := range transactions {
		wg.Add(1)
		txChan <- txID
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	txCount := 0
	succTxCount := 0
	failedTxCount := 0

	for result := range resultChan {
		txCount += 1
		if result.err != nil {
			fmt.Println(fmt.Sprintf("Got an error while processing tx id '%d': %v", result.txID, result.err))
			failedTxCount += 1
		} else {
			succTxCount += 1
		}
	}

	fmt.Println("Total transactions: ", txCount)
	fmt.Println("Successful transactions: ", succTxCount)
	fmt.Println("Failed transactions: ", failedTxCount)

}

func worker(txChan <-chan int, resultChan chan<- txResult, wg *sync.WaitGroup) {
	for txID := range txChan {
		func() {
			defer func() {
				// Note: if needed I can handle panics here
				wg.Done()
			}()
			success, err := processTransaction(txID)
			resultChan <- txResult{
				txID:    txID,
				success: success,
				err:     err,
			}
		}()
	}
}

func processTransaction(id int) (bool, error) {
	// Simulate random success or failure
	if id%3 == 0 { // Fail every 3rd transaction
		return false, fmt.Errorf("transaction %d failed", id)
	}
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100))) // Simulate processing time
	return true, nil
}
