package main

import (
	"time"
)

type Product struct {
	ID        int       `bson:"_id, omitempty"`
	Name      string    `bson:"name"`
	Category  string    `bson:"category"`
	Price     float64   `bson:"price"`
	Stock     int       `bson:"stock"`
	Tags      []string  `bson:"tags"`
	CreatedAt time.Time `bson:"created_at"`
}
