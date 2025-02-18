package main

import "fmt"

func main() {
	// Asignación simple
	var a string = "Hola"
	fmt.Println("a:", a)

	// Asignación con suma
	b := 10
	b += 5 // Ahora b es 15 b = b + 5
	fmt.Println("b:", b)

	// Asignación con resta
	c := 10
	c -= 5 // Ahora c es 5 c = c - 5
	fmt.Println("c:", c)

	// Asignación con multiplicación
	d := 10
	d *= 5 // Ahora d es 50 d = d * 5
	fmt.Println("d:", d)

	// Asignación con división
	e := 10
	e /= 5 // Ahora e es 2 e = e / 5
	fmt.Println("e:", e)

	// Asignación con módulo
	f := 13
	f %= 5 // Ahora f es 3 f = f % 5 => 13 % 5 = 3
	fmt.Println("f:", f)

}
