package db

import (
	"context"
	"fmt"
	"log"

	//"reflect"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func DB_connect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	DB = client.Database("shortenurl")
	fmt.Println("Connected to MongoDB!")
}

func GetCollection(name string) *mongo.Collection {
	Col := DB.Collection(name)
	return Col
}
