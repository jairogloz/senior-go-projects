package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var x int32
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&x, 1)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("(main)x: ", x)

}
