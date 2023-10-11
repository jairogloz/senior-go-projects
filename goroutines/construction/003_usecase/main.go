package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	begin := time.Now()

	wg.Add(3)

	go func() {
		defer wg.Done()
		res, err := http.Get("https://httpbin.org/delay/3")
		if err != nil {
			fmt.Printf("Error in request: %v\n", err)
			return
		}
		defer res.Body.Close()

		body, _ := io.ReadAll(res.Body)
		fmt.Printf("Response 1: %s\n", string(body))
	}()

	go func() {
		defer wg.Done()
		res, err := http.Get("https://httpbin.org/delay/3")
		if err != nil {
			fmt.Printf("Error in request: %v\n", err)
			return
		}
		defer res.Body.Close()

		body, _ := io.ReadAll(res.Body)
		fmt.Printf("Response 2: %s\n", string(body))
	}()

	go func() {
		defer wg.Done()
		res, err := http.Get("https://httpbin.org/delay/4")
		if err != nil {
			fmt.Printf("Error in request: %v\n", err)
			return
		}
		defer res.Body.Close()

		body, _ := io.ReadAll(res.Body)
		fmt.Printf("Response 3: %s\n", string(body))
	}()

	wg.Wait()

	fmt.Printf("All requests completed in %.2f seconds\n", time.Since(begin).Seconds())
}
