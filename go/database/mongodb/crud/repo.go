package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var ctx context.Context
var client *mongo.Client
var collection *mongo.Collection

func CloseConnection() {
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()
}

func SetupDB() {
	var err error

	ctx = context.TODO()
	client, err = mongo.Connect(options.Client().ApplyURI("mongodb://mongoadmin:mongoadmin@localhost:27017/"))

	if err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	collection = client.Database("mydb").Collection("customers")

	err = collection.Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
