package main

import "fmt"

func main() {
	// Asignación simple
	var a string = "Hola"
	fmt.Println(a)

	// Asignación con suma
	b := 10
	b += 5 // Ahora b es 15
	fmt.Println(b)

	// Asignación con resta
	c := 10
	c -= 5 // Ahora c es 5
	fmt.Println(c)

	// Asignación con multiplicación
	d := 10
	d *= 5 // Ahora d es 50
	fmt.Println(d)

	// Asignación con división
	e := 10
	e /= 5 // Ahora e es 2
	fmt.Println(e)

	// Asignación con módulo
	f := 13
	f %= 5 // Ahora f es 3
	fmt.Println(f)

}
