package main

import (
	"encoding/json"
	"fmt"
	domain "github.com/JairoGLoz/senior-go-projects/senior-go-projects/restful_api/000_domain"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/users/", userHandler)

	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// auth
	fmt.Printf("Received request method: %s\n", r.Method)
	switch r.Method {
	case http.MethodGet:
		handleGetUser(w, r)
	case http.MethodPost:
		handleCreateUser(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	// Split the URL path by '/' and get the last part, which should be the id
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	// Here you can add the logic to get the user by id
	// For now, let's just return the id in the response
	fmt.Fprintf(w, "Requested user with id: %s", id)
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
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
