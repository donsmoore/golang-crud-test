package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func main() {

	// connect db
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	// get table
	collection := client.Database("testing").Collection("numbers")

	// insert a row
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, _ := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})

	// show new row id
	id := res.InsertedID
	fmt.Println("Done: ",  id , " success.")

}
