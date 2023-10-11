package main

import "fmt"

func main() {

	bandaTransp := make(chan string) // capacidad 0

	go func() {
		fmt.Println("enviando pieza 1 a banda...")
		bandaTransp <- "pieza 1"
		fmt.Println("pieza 1 enviada")
	}()

	fmt.Println("esperando pieza 1 en banda...")
	pieza1 := <-bandaTransp
	fmt.Println("pieza 1 recibida: ", pieza1)

}
