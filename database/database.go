package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//GetMongoDbConnection get connection of mongodb
func getMongoDbConnection() (*mongo.Client, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://127.0.0.1:27017"))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}

//GetMongoDbCollection get specific collection from the database
func GetMongoDbCollection(CollectionName string) (*mongo.Collection, error) {
	client, err := getMongoDbConnection()

	if err != nil {
		return nil, err
	}

	collection := client.Database("goapi").Collection(CollectionName)

	return collection, nil
}
