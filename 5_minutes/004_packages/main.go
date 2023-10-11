package main

import (
	"fmt"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/5_minutes/004_packages/my_package"
	"math"
)

func main() {
	fmt.Println("Hello from main")
	fmt.Println(math.Max(7, 8.5))

	my_package.MyFunc()
}
