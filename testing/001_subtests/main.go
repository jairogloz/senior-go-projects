package main

import (
	"fmt"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/testing/000_basic_testing/core"
)

func main() {

	e := core.Employee{
		Name:  "Jairo",
		Email: "emaildomain.com",
		Age:   18,
	}

	isValid, err := core.Validate(e)
	fmt.Println()
	fmt.Println("Is valid: ", isValid)
	fmt.Println("Error: ", err)

}
