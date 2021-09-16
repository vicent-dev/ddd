package ddd

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnection struct {
	cnn *mongo.Client
}

func NewMongoConnection(user, password, host, port string) *MongoConnection {

	clientOpts := options.Client().ApplyURI(fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/?authSource=admin",
		user,
		password,
		host,
		port,
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

func (mc *MongoConnection) GetDatabase(dbName string) *mongo.Database {
	return mc.cnn.Database(dbName)
}
