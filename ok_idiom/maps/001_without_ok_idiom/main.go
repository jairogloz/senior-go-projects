package main

import "fmt"

func main() {

	var myVar interface{}
	// type = interface{}
	// value = 1.5 (float64)
	myVar = 1.5

	// Necesitamos tratar a myVar como int
	// para poder sumarle 1
	var sum int
	a := 1
	sum = a + myVar.(int)

	fmt.Println("sum: ", sum)
}
