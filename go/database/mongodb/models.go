package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID        primitive.ObjectID `bson:"_id, omitempty"`
	Name      string             `bson:"name"`
	Category  string             `bson:"category"`
	Price     float64            `bson:"price"`
	Stock     int                `bson:"stock"`
	Tags      []string           `bson:"tags"`
	CreatedAt time.Time          `bson:"created_at"`
}
