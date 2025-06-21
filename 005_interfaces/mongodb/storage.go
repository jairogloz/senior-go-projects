package mongodb

import (
	"context"
	"github.com/JairoGLoz/senior-go-projects/senior-go-projects/005_interfaces/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Storage struct {
	client *mongo.Client
}

func NewStorage() core.EmployeeStorage {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	// Ping the MongoDB server to verify that the connection is working
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	return &Storage{client: client}
}

func (s Storage) GetEmployee(id string) (*core.Employee, error) {
	collection := s.client.Database("mydb").Collection("employees")

	filter := bson.M{"_id": id}

	// Find the user document that matches the filter
	var result core.Employee
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s Storage) SaveEmployee(employee *core.Employee) error {
	// Get a handle for the employees collection
	employeesCollection := s.client.Database("testdb").Collection("employees")

	// Insert a new employee document
	_, err := employeesCollection.InsertOne(context.Background(), employee)
	if err != nil {
		return err
	}

	return nil
}
