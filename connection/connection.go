package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Product *mongo.Collection
var City *mongo.Collection
var Province *mongo.Collection
var Customer *mongo.Collection
var ShippingMethod *mongo.Collection
var Basket *mongo.Collection

func init() {
	godotenv.Load()
	mongoURI := os.Getenv("MONGO_URI")
	mongoDB := os.Getenv("MONGO_DB")
	client, conErr := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if conErr != nil {
		fmt.Println("error in db connection")
		os.Exit(1)
	}
	db := client.Database(mongoDB)
	Product = db.Collection("products")
	City = db.Collection("cities")
	Province = db.Collection("provinces")
	Customer = db.Collection("customers")
	ShippingMethod = db.Collection("shippingMethods")
	Basket = db.Collection("baskets")
}

// for test purpose
func init() {
	count, _ := Customer.CountDocuments(context.Background(), bson.M{})
	if count == 0 {
		Customer.InsertOne(context.TODO(), bson.M{"phone": "09107655173"})
	}
}
