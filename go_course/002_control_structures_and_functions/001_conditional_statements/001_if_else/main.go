package main

import "fmt"

func main() {
	// if // si
	// else // sino
	// else if

	// Simple if
	//if temperatura > 30 { // Si la temperatura es mayor a 30
	//	fmt.Println("Hace calor afuera")
	//}

	// if else
	//if temperatura > 30 { // Si la temperatura es mayor a 30
	//	fmt.Println("Hace calor afuera")
	//} else { // Si no (si la temperatura es menor o igual a 30)
	//	fmt.Println("La temperatura es agradable")
	//}

	temperatura := 9

	// if else if
	if temperatura > 30 {
		fmt.Println("Hace calor afuera")
	} else if temperatura < 10 {
		fmt.Println("Hace frío afuera")
	} else if temperatura < 20 {
		fmt.Println("hace mucho frío")
	} else {
		fmt.Println("La temperatura es agradable")
	}

}
