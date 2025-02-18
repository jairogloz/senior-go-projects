package package_a

import "fmt"

var ENV = "dev"

func init() {
	fmt.Println("init in package_a/file_a.go")
	fmt.Println("ENV:", ENV)
}

func FunctionA() {
	fmt.Println("FunctionA in package_a/file_a.go")
}
