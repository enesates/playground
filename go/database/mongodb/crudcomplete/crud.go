package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Category  string             `bson:"category"`
	Price     float64            `bson:"price"`
	Stock     int                `bson:"stock"`
	Tags      []string           `bson:"tags"`
	CreatedAt time.Time          `bson:"created_at"`
}

var collection *mongo.Collection

func connectDB() (*mongo.Client, error) {
	// Connection String (für lokales MongoDB oder MongoDB Atlas)
	uri := "mongodb://mongoadmin:mongoadmin@localhost:27017"

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping testen
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Erfolgreich mit MongoDB verbunden!")
	return client, nil
}

func InsertProduct(product Product) (primitive.ObjectID, error) {
	product.CreatedAt = time.Now()

	result, err := collection.InsertOne(context.TODO(), product)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func InsertManyProducts(products []Product) ([]primitive.ObjectID, error) {
	// Interface slice für InsertMany erstellen
	docs := make([]interface{}, len(products))
	for i, p := range products {
		p.CreatedAt = time.Now()
		docs[i] = p
	}

	result, err := collection.InsertMany(context.TODO(), docs)
	if err != nil {
		return nil, err
	}

	// IDs konvertieren
	ids := make([]primitive.ObjectID, len(result.InsertedIDs))
	for i, id := range result.InsertedIDs {
		ids[i] = id.(primitive.ObjectID)
	}

	return ids, nil
}

func FindProductByID(id primitive.ObjectID) (*Product, error) {
	var product Product
	filter := bson.M{"_id": id}

	err := collection.FindOne(context.TODO(), filter).Decode(&product)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("produkt mit ID %s nicht gefunden", id.Hex())
	}
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func FindProductsByCategory(category string) ([]Product, error) {
	filter := bson.M{"category": category}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var products []Product
	for cursor.Next(context.TODO()) {
		var product Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func UpdateProductPrice(id primitive.ObjectID, newPrice float64) error {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"price": newPrice}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("kein Produkt mit ID %s gefunden", id.Hex())
	}

	fmt.Printf("Preis aktualisiert: %d Dokumente geändert\n", result.ModifiedCount)
	return nil
}

func DeleteProductByID(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("kein Produkt mit ID %s gefunden", id.Hex())
	}

	fmt.Printf("Produkt gelöscht: %d Dokumente\n", result.DeletedCount)
	return nil
}

func CreateIndexes() error {
	// Index auf Category
	categoryIndex := mongo.IndexModel{
		Keys: bson.D{{Key: "category", Value: 1}},
	}

	// Index auf Price
	priceIndex := mongo.IndexModel{
		Keys: bson.D{{Key: "price", Value: 1}},
	}

	_, err := collection.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		categoryIndex,
		priceIndex,
	})

	if err != nil {
		return err
	}

	fmt.Println("Indexes erstellt!")
	return nil
}

func main() {
	client, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := client.Disconnect(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
	}()

	db := client.Database("testdb")
	collection = db.Collection("products")

	// Collection leeren für sauberen Test
	err = collection.Drop(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Indexes erstellen
	//err = CreateIndexes()
	//if err != nil {
	//	log.Fatal(err)
	//}

	// Test: Ein Produkt einfügen
	product := Product{
		Name:     "Laptop",
		Category: "Electronics",
		Price:    999.99,
		Stock:    10,
		Tags:     []string{"computer", "portable"},
	}

	id, err := InsertProduct(product)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Produkt eingefügt mit ID: %s\n", id.Hex())

	// Test: Mehrere Produkte einfügen
	products := []Product{
		{Name: "Mouse", Category: "Electronics", Price: 29.99, Stock: 50, Tags: []string{"computer", "accessory"}},
		{Name: "Keyboard", Category: "Electronics", Price: 79.99, Stock: 30, Tags: []string{"computer", "accessory"}},
		{Name: "Go Programming", Category: "Books", Price: 39.99, Stock: 100, Tags: []string{"programming", "education"}},
		{Name: "T-Shirt", Category: "Clothing", Price: 19.99, Stock: 200, Tags: []string{"casual"}},
	}

	ids, err := InsertManyProducts(products)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d Produkte eingefügt\n", len(ids))

	// Test: Produkt nach ID finden
	foundProduct, err := FindProductByID(id)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Printf("\nProdukt gefunden: %+v\n", foundProduct)
	}

	// Test: Produkte nach Kategorie finden
	electronics, err := FindProductsByCategory("Electronics")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%d Electronics gefunden:\n", len(electronics))
	for _, p := range electronics {
		fmt.Printf("- %s: %.2f€\n", p.Name, p.Price)
	}

	// Test: Preis aktualisieren
	err = UpdateProductPrice(id, 899.99)
	if err != nil {
		log.Println(err)
	}

	// Test: Produkt löschen
	err = DeleteProductByID(ids[0])
	if err != nil {
		log.Println(err)
	}
}
