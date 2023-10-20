package main

import (
	"fmt"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/testing/003_mocks_self_made/employee_service"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/testing/003_mocks_self_made/employee_storage"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		panic("MONGO_URI is not set")
	}

	mongoClient, disconnectFunc, err := employee_storage.NewMongoClient(mongoURI)
	if err != nil {
		panic(err)
	}
	defer disconnectFunc()

	employeeStorage := employee_storage.NewEmployeeStorage(mongoClient)

	employeeService := employee_service.NewService(employeeStorage)

	ids := []string{
		"64d7b8698029c0ab12a2cd1f",  // valid
		"64d7b8698029c0ab12a2cd1e",  // invalid
		"64d7b8698029c0ab12a2cd1a",  // not found
		"64d7b8698029c0ab12a2cd1g_", // error converting id to object id
	}

	for _, id := range ids {
		isValid, err := employeeService.ValidateEmployee(id)
		fmt.Printf("\nid: %s, isValid: %v, err: %v\n\n", id, isValid, err)
	}

	fmt.Printf(fmt.Sprintf("\n\n\t\t**************************\n\t\t*  Suscríbete al canal!  *\n\t\t**************************\n\n"))

}
