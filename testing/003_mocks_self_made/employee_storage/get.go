package employee_storage

import (
	"context"
	"fmt"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/testing/003_mocks_self_made/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s Storage) GetEmployee(id string) (*domain.Employee, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("error converting id to object id: %w", err)
	}

	filter := bson.M{"_id": objectID}

	var result domain.Employee
	err = s.Col.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrNotFound
		}
		return nil, fmt.Errorf("query failed: %w", err)
	}

	return &result, nil
}
