package main

import (
	"encoding/json"
	"fmt"
	domain "github.com/JairoGLoz/senior-go-projects/senior-go-projects/restful_api/000_domain"
	"net/http"
)

func main() {
	http.HandleFunc("/create-user", createUserHandler)

	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {

	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	if user.Age < 18 {
		http.Error(w, "User must be 18 years or older", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "User created successfully!")
}
