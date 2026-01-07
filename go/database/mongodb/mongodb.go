package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Customer struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func GetCustomer(ctx context.Context, client *mongo.Client) {
	collection := client.Database("mydb").Collection("customers")

	filter := bson.D{{"name", "Customer"}}
	var customer Customer
	err := collection.FindOne(ctx, filter).Decode(&customer)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(customer)
}

func AddCustomer(ctx context.Context, client *mongo.Client) {
	collection := client.Database("mydb").Collection("customers")
	_, err := collection.InsertOne(ctx, Customer{"Customer", 40})

	if err != nil {
		log.Fatal(err)
	}
}

func UpdateCustomer(ctx context.Context, client *mongo.Client) {
	collection := client.Database("mydb").Collection("customers")
	filter := bson.D{{"name", "Customer"}}

	update := bson.D{{"$set", bson.D{{"age", 50}}}}
	_, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ctx := context.TODO()
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://mongoadmin:mongoadmin@localhost:27017/?authSource=admin"))

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Disconnect(ctx); err != nil {

			panic(err)
		}
	}()

	if err = client.Ping(ctx, nil); err != nil {

		log.Fatal(err)
	}

	AddCustomer(ctx, client)
	GetCustomer(ctx, client)
	UpdateCustomer(ctx, client)
	GetCustomer(ctx, client)
}
