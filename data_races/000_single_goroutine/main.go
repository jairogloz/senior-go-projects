package main

import (
	"fmt"
)

func main() {

	x := 0

	for i := 0; i < 1000; i++ {
		x++
	}

	fmt.Println("x: ", x)

}
