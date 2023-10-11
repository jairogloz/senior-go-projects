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

	myInt, ok := myVar.(int)
	if ok == true { // type assertion tuvo exito
		sum = a + myInt
		fmt.Println("sum", sum)
	} else { // type assertion no tuvo exito
		fmt.Printf("myVar no es de tipo int, es type %T\n", myVar)
		fmt.Println("myInt", myInt)
	}
}
