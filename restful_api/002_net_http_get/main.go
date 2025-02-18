package main

import (
	"encoding/json"
	"fmt"
	domain "github.com/JairoGLoz/senior-go-projects/senior-go-projects/restful_api/000_domain"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/create-user", createUserHandler)
	http.HandleFunc("/get-user/", getUserHandler)

	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Split the URL path by '/' and get the last part, which should be the id
	// ["", "get-user", "1"]
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	// Here you can add the logic to get the user by id
	// For now, let's just return the id in the response
	fmt.Fprintf(w, "Requested user with id: %s", id)
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
