package main

import (
	"encoding/json"
	"fmt"
	domain "github.com/JairoGLoz/senior-go-projects/senior-go-projects/restful_api/000_domain"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

func main() {
	router := httprouter.New()
	router.GET("/users/:id", authMiddleware(handleGetUser))
	router.POST("/users/", authMiddleware(handleCreateUser))

	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}

func authMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Pass the request to the next handler if the Authorization header is valid
		next(w, r, ps)
	}
}

func handleGetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	fmt.Fprintf(w, "Requested user with id: %s", id)
}

func handleCreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User created successfully!")
}
