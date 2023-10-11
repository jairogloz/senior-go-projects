package main

import (
	"fmt"
	"time"
)

func main() {

	go myFunc()

	fmt.Println("Waiting from main")
	time.Sleep(3 * time.Second)

	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Hello from channel"
	}()

	fmt.Println("Waiting from value from channel")
	fmt.Println("Value from channel: ", <-ch)
}

func myFunc() {
	time.Sleep(2 * time.Second)
	fmt.Println("Hello from myFunc")
}
