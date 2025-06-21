package http

import (
	"encoding/json"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/006_cybersecurity/internal/domain"
	"net/http"
	"strconv"
)

// CreateUserHandler creates a new user
func CreateUserHandler(service domain.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*var input dto.UserCreateParams

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&input); err != nil {
			http.Error(w, fmt.Sprintf("invalid request body: %s", err.Error()), http.StatusBadRequest)
			return
		}
		//disable unknwon fields

		// Validate input (if using a validation library, e.g., go-playground/validator)
		v := validator.New()
		if err := v.Struct(input); err != nil {
			http.Error(w, "validation failed", http.StatusBadRequest)
			return
		}

		// Map DTO to domain.User
		now := time.Now()
		user := domain.User{
			Name:      input.Name,
			Age:       input.Age,
			CreatedAt: &now,
			IsAdmin:   false, // Default value or based on additional logic
		}*/

		var user domain.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		user, err := service.Create(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

// GetUserHandler retrieves a user by ID
func GetUserHandler(service domain.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid id parameter", http.StatusBadRequest)
			return
		}

		user, err := service.GetByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
