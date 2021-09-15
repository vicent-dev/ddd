package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnection struct {
	cnn *mongo.Client
}

func NewMongoConnection() *MongoConnection {

	clientOpts := options.Client().ApplyURI(fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/?authSource=admin",
		os.Getenv("MONGO_USERNAME"),
		os.Getenv("MONGO_PASSWORD"),
		os.Getenv("MONGO_HOST"),
		os.Getenv("MONGO_PORT"),
	))

	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return &MongoConnection{client}
}

func (mc *MongoConnection) GetDatabase() *mongo.Database {
	return mc.cnn.Database(os.Getenv("MONGO_DATABASE"))
}
