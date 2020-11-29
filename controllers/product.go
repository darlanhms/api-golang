package controllers

import (
	"api-golang/database"
	"api-golang/models"
	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateProduct register a product in the database
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

// GetProducts list all the products that are registered in the database
func GetProducts() ([]byte, error) {
	productsCollection, err := database.GetMongoDbCollection("products")

	if err != nil {
		return nil, err
	}

	var results []bson.M

	cursor, err := productsCollection.Find(context.Background(), bson.M{})

	defer cursor.Close(context.Background())

	if err != nil {
		return nil, err
	}

	cursor.All(context.Background(), &results)

	products, _ := json.Marshal(results)

	return products, nil
}

// DeleteProductByID delete a product with specific ID
func DeleteProductByID(id string) (*mongo.DeleteResult, error) {
	productsCollection, err := database.GetMongoDbCollection("products")

	if err != nil {
		return nil, err
	}

	objID, _ := primitive.ObjectIDFromHex(id)

	result, err := productsCollection.DeleteOne(context.Background(), bson.M{"_id": objID})

	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateProduct update a product that has the id passed throught the param
func UpdateProduct(id string, product models.Product) (*mongo.UpdateResult, error) {
	productsCollection, err := database.GetMongoDbCollection("products")

	if err != nil {
		return nil, err
	}

	objID, _ := primitive.ObjectIDFromHex(id)

	update := bson.M{
		"$set": product,
	}

	result, err := productsCollection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)

	if err != nil {
		return nil, err
	}

	return result, nil

}
