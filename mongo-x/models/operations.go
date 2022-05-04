package models

import (
	"context"
	"example/mongo/connection"
	"fmt"
)

func AddUser() {
	client, err := connection.Create()

	coll := client.Database("users").Collection("users")
	if err != nil {
		panic(err)
	}

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
