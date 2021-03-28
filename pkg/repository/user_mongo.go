package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoUserRepository gives mongo implementation for user operations
type MongoUserRepository struct {
	DB *mongo.Database
}

// AddUser adds new user to the system
func (usr *MongoUserRepository) AddUser(user UserEntity) (string, error) {
	result, err := usr.DB.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		return "", err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}
