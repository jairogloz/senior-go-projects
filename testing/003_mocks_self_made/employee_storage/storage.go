package employee_storage

import (
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/testing/003_mocks_self_made/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

// Storage implements the domain.EmployeeStorage interface
type Storage struct {
	Col *mongo.Collection
}

func NewEmployeeStorage(mongoClient *mongo.Client) domain.EmployeeStorage {
	return &Storage{
		Col: mongoClient.Database("hr").Collection("employees"),
	}
}
