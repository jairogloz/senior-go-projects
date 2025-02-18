package my_math_test

import (
	"fmt"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/001_function_comments/my_math"
)

func ExampleAdd() {
	sum := my_math.Add(2, 3)
	fmt.Println(sum)
	// Output: 5
}
