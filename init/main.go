package main

import (
	"fmt"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/init/package_a"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/init/package_b"
)

func init() {
	fmt.Println("init in main.go")

}

func main() {
	fmt.Println("main in main.go")

	package_a.ENV = "prod"

	package_a.FunctionA()

	package_b.FunctionB()

}
