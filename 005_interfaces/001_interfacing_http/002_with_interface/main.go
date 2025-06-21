package main

import (
	"fmt"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/005_interfaces/001_interfacing_http/002_with_interface/internal/services/cat"
)

// CatAPIHandler defines the interface with the GetCatFacts method
type CatAPIHandler interface {
	GetCatFacts(n int) ([]string, error)
}

func main() {

	var handler CatAPIHandler
	handler = &cat.APIHandler{
		Url: "https://meowfacts.herokuapp.com/?count=%d",
	}

	// Call with n = 1
	facts, err := handler.GetCatFacts(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Cat Facts (n=1):", facts)
	}

	// Call with n = 3
	facts, err = handler.GetCatFacts(3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Cat Facts (n=3):", facts)
	}

	// Call with n = -1
	facts, err = handler.GetCatFacts(-1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Cat Facts (n=-1):", facts)
	}
}
