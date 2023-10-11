package main

import (
	"fmt"
	"sync"
)

func main() {

	x := 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			x++
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("(main)x: ", x)

}
