package package_b

import (
	"fmt"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/init/package_a"
)

func init() {
	fmt.Println("init in package_b/file_b.go")
	fmt.Println("ENV:", package_a.ENV)
}

func FunctionB() {
	fmt.Println("FunctionB in package_b/file_b.go")
}
