package main

import (
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetProductByID(id interface{}) {
	var product Product

	filter := bson.D{{"_id", id}}

	err := collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(product)
}

func AddProduct(product Product) interface{} {
	res, err := collection.InsertOne(ctx, product)

	if err != nil {
		log.Fatal(err)
	}

	return res.InsertedID
}

func GetProductByCategory(category string) {
	println("\nGetProductByCategory\n-----------")

	filter := bson.D{{"category", category}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product Product
		err := cursor.Decode(&product)

		if err != nil {
			log.Fatal(err)
			return
		}

		GetProductByID(product.ID)
	}
}
