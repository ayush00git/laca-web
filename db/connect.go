package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo (uri, dbName string) (*mongo.Database, error) {
	opts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}


	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return client.Database(dbName), nil
}
