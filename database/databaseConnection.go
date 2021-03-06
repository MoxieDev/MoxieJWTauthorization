package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
)

//DBinstance func
func DBinstance() *mongo.Client {
	err := godotenv.Load("variables.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MongoDB := os.Getenv("MONGODB_URL")
	fmt.Println(MongoDB)

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDB))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	return client
}

//Client Database Instance
var Client *mongo.Client = DBinstance()

//OpenCollection is a function that makes a connection with a collection in the database
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)

	return collection
}
