package controllers

import (
	"api-golang/database"
	"api-golang/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateProduct(product models.Product) (*mongo.InsertOneResult, error) {
	productsCollection, err := database.GetMongoDbCollection("products")

	if err != nil {
		return nil, err
	}

	result, err := productsCollection.InsertOne(context.Background(), product)

	if err != nil {
		return nil, err
	}

	return result, nil
}
