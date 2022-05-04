package main

import (
	"context"
	"encoding/json"
	"example/mongo/connection"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	client, err := connection.Create()
	if err != nil {

		panic(err)
	}
	coll := client.Database("users").Collection("users")

	name := "Kai Havertz"
	var result bson.M

	err = coll.FindOne(context.TODO(), bson.D{{Key: "name", Value: name}}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document with %s was found", name)
		return
	}

	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}
