package main

import (
	"fmt"
	http2 "github.com/JairoGLoz/senior-go-projects/senior-go-projects/006_cybersecurity/internal/adapters/http"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/006_cybersecurity/internal/adapters/storage"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/006_cybersecurity/internal/application"
	"net/http"
)

func main() {
	userStorage := storage.NewInMemoryUserStorage()
	service := application.NewUserService(userStorage)

	http.HandleFunc("/create", http2.CreateUserHandler(service))
	http.HandleFunc("/get", http2.GetUserHandler(service))

	fmt.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}
}
