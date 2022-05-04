package models

import (
	"context"
	"example/mongo/connection"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client, err = connection.Create()

func init() {
	if err != nil {
		panic(err)
	}
}

func CreateUser() {

	coll := client.Database("users").Collection("users")

	user := &User{
		Name:    "Bill Greatness",
		Age:     45,
		Country: "Ghana",
	}

	if err = user.Validate(); err != nil {
		panic(err)
	}

	result, err := coll.InsertOne(context.TODO(), user)

	if err != nil {
		panic(err)
	}

	fmt.Print(result)
}

func UpdateUser() {

	coll := client.Database("users").Collection("users")

	userData := &User{
		Name: "Bill Greatnes$",
	}

	filter := bson.D{{Key: "name", Value: "Bill Greatness"}}
	_, err = coll.UpdateOne(context.TODO(), filter, userData)

	if err != nil {
		panic(err)
	}

}

func DeleteUsers() {
	coll := client.Database("users").Collection("users")

	filter := bson.D{{Key: "age", Value: bson.D{{Key: "$gt", Value: 90}}}}

	_, err := coll.DeleteMany(context.TODO(), filter)

	if err != nil {
		panic(err)
	}
}

func ReadUsers() *mongo.Cursor {

	coll := client.Database("users").Collection("users")
	re := "/[a-z]\\d+@apple.com/i"

	filter := bson.D{{Key: "email", Value: bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: re}}}}}
	// try reading all users with an apple domain extension.

	cursor, err := coll.Find(context.TODO(), filter)

	if err != nil {
		panic(err)
	}

	return cursor

}
