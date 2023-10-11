package main

import "fmt"

func main() {

	var myInt int = 5
	var myString string
	myFloat := 5.0
	myBool := true

	if myInt > 4 {
		fmt.Println("myInt is greater than 4")
	} else {
		fmt.Println("myInt is not greater than 4")
	}

	result := myFunc(myString, myFloat, myBool)

	fmt.Println(result)
}

func myFunc(s string, f float64, b bool) string {
	return fmt.Sprintf("myString: %s, myFloat: %f, myBool: %t", s, f, b)
}
