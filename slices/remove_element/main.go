package main

import (
	"fmt"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/slices/remove_element/util"
)

func main() {

	x := []int{10, 20, 30, 40, 50, 60, 70, 80}

	fmt.Println("x[1:3]: ", x[1:3])
	fmt.Println("x[:5]: ", x[:5])
	fmt.Println("x[2:]: ", x[2:])

	fmt.Println("x: ", x)
	x = util.RemoveLoop(x, 3)
	fmt.Println("x: ", x)
	x = util.RemoveNewSlice(x, 3)
	fmt.Println("x: ", x)
	x = util.RemoveCopy(x, 3)
	fmt.Println("x: ", x)
	x = util.RemoveAppend(x, 3)
	fmt.Println("x: ", x)
	x = util.RemoveWithSlicesPackage(x, 3)
	fmt.Println("x: ", x)

}
