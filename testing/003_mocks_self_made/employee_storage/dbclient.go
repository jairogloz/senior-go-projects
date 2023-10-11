package employee_storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

func NewMongoClient(mongoUri string) (client *mongo.Client, disconnectClientFunc func(), err error) {

	opts := options.Client().SetMaxPoolSize(2).
		ApplyURI(mongoUri)

	client, err = mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Println("error connecting mongo client", "error", err.Error())
		return nil, nil, fmt.Errorf("error connecting mongo client: %s", err.Error())
	}

	disconnectFunc := func() {
		err := client.Disconnect(context.Background())
		if err != nil {
			log.Println("error disconnecting mongo client", "error", err.Error())
		}
		log.Println("mongo client has been successfully disconnected")
	}

	err = client.Ping(context.Background(), readpref.Secondary())
	if err != nil {
		log.Println("error pinging mongo client", "error", err.Error())
		return nil, nil, err
	}

	return client, disconnectFunc, nil

}
