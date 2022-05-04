package main

import (
	"context"
	"example/mongo/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	var info []bson.M
	result := models.ReadUsers()

	if err := result.All(context.TODO(), &info); err != nil {
		panic(err)
	}

	fmt.Print(info)

	// models.DeleteUsers()

}
