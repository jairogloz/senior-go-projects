package main

import (
	"encoding/json"
	"fmt"
	domain "github.com/JairoGLoz/senior-go-projects/senior-go-projects/restful_api/000_domain"
	"net/http"
	"strings"
)

func main() {
	http.Handle("/users/", authMiddleware(http.HandlerFunc(userHandler)))

	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Pass the request to the next handler if the Authorization header is valid
		next.ServeHTTP(w, r)
	})
}

func userHandler(w http.ResponseWriter, r *http.Request) {

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
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]
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
