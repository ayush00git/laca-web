package helpers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

func CodeIndex(db *mongo.Database) error {
	collection := db.Collection("courses")

	_, err := collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{{Key: "code", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	return err
}

func StudentIndex(db *mongo.Database) error {
	collection := db.Collection("students")

	_, err := collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{{Key: "roll_number", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	})
	return err
}
