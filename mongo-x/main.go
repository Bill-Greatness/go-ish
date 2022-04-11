package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// getting started with the mongoDB driver.

	// check availability with godotenv

	if err := godotenv.Load(); err != nil {
		log.Println("Go env file is not available.")
	}

	// get uri from environment.
	uri := os.Getenv("MONGO_URI")

	// check if uri is not empty or exit with a message.
	if uri == "" {
		log.Fatal("You must set MONGO_URI environment variable")
	}

	// connecting to the client.
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	// check if there is connection with the connection uri above
	if err != nil {
		panic(err)
	}

	// try disconnection the cotext.TODO from client.
	// inline function.
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// access content.
	coll := client.Database("test").Collection("student")
	title := "Hello Greatness"

	var result bson.M

	err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		log.Println("There are no documents in this collection.")
	}

	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Your query data :%v", jsonData)
}
