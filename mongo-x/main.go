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
		log.Print("MONGO Environment Variable Not Set")
	}

	uri := os.Getenv("MONGO_URI")

	// check if uri is empty

	if uri == "" {
		log.Fatal("You must set your MONGO_URI environment variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}
	// check and disconnect after everything
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// get column data.
	coll := client.Database("test").Collection("students")

	title := "Hello World"

	var result bson.M

	err = coll.FindOne(context.TODO(), bson.D{{Key: "title", Value: title}}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document with title %s was found", title)
		return
	}

	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", jsonData)

	// finding multiple documents
	filter := bson.D{{Key: "pop", Value: bson.D{{Key: "$lte", Value: 500}}}}

	cursor, err := coll.Find(context.TODO(), filter)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", cursor)

}
