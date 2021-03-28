package datastore

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoStore provides mongodb datastore
type MongoStore struct{}

// Init initializes connection to mongodb server
func Init() (*mongo.Database, error) {
	val, exist := os.LookupEnv("DATABASE_URL")
	if !exist {
		return nil, errors.New("DATABASE_URL value not set")
	}

	db, exist := os.LookupEnv("DATABASE_NAME")
	if !exist {
		return nil, errors.New("DATABASE_NAME value not set")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(val))
	if err != nil {
		return nil, err
	}

	return client.Database(db), nil
}
