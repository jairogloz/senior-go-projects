package main

import "fmt"

type employee struct {
	Name string
	Age  int
}

func main() {
	mySlice := []int{1, 2, 3, 4, 5}

	fmt.Println(mySlice[0])
	fmt.Println(mySlice[2])

	e1 := employee{
		Name: "John",
		Age:  30,
	}

	myMap := map[string]employee{
		"001": e1,
	}

	fmt.Println(myMap["001"].Name)
}
