package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	coll := client.Database("users").Collection("users")

	name := "Kai Havertz"

	var result bson.M

	err = coll.FindOne(context.TODO(), bson.D{{Key: "name", Value: name}}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document with Name: %s was found", name)
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

	// inserting a document.

	doc := bson.D{{Key: "title", Value: "Beyond Intelligence"}, {Key: "content", Value: "We can make things right!"}}

	_, err = coll.InsertOne(context.TODO(), doc)

	if err != nil {
		panic(err)
	}

	// inserting multiple documents.
	docs := []interface{}{
		bson.D{{Key: "name", Value: "Bill Greatness"}, {Key: "status", Value: "Active"}},
		bson.D{{Key: "name", Value: "Anon"}, {Key: "status", Value: "Active"}},
	}

	_, err = coll.InsertMany(context.TODO(), docs)
	if err != nil {
		panic(err)
	}

	// updating a document.
	// get ID field that needs the update
	id, _ := primitive.ObjectIDFromHex("someRandomeID")

	filter = bson.D{{Key: "_id", Value: id}}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "rating", Value: 4.42}}}}

	_, err = coll.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		panic(err)
	}

	//updating multiple documents.

	filter = bson.D{{Key: "address.market", Value: "Syndey"}}

	update = bson.D{{Key: "$mul", Value: bson.D{{Key: "price", Value: 1.22}}}}

	_, err = coll.UpdateMany(context.TODO(), filter, update)

	if err != nil {
		panic(err)
	}

	// delete a document.
	filter = bson.D{{Key: "name", Value: "Greatness"}}
	_, err = coll.DeleteOne(context.TODO(), filter)

	if err != nil {
		panic(err)
	}
}
