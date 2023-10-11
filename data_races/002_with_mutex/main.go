package main

import (
	"fmt"
	"sync"
)

func main() {

	x := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			x++
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("(main)x: ", x)

}
