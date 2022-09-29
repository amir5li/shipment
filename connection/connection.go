package connection

import (
	"context"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var City *mongo.Collection
var Province *mongo.Collection
func init(){
	godotenv.Load()
	mongoURI := os.Getenv("MONGO_URI")
	mongoDB := os.Getenv("MONGO_DB")
	client , conErr := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if conErr != nil {
		fmt.Println("error in db connection")
		os.Exit(1)
	}
	db := client.Database(mongoDB)
	City = db.Collection("cities")
	Province = db.Collection("provinces")
}