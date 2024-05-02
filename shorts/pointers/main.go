package main

import "fmt"

func main() {
	var a string = "hola"
	fmt.Println(a)

	var pA *string = &a
	fmt.Println(pA)
	fmt.Println(*pA)
	*pA = "adios"
	fmt.Println(a)
}
