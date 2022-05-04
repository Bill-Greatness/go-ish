package models

import (
	"github.com/go-playground/validator/v10"
)

// user model definitions.
type User struct {
	Name    string `bson:"name,omitempty" validate:"required"`
	Age     int32  `bson:"age,omitempty" validate:"required"`
	Email   string `bson:"email"`
	Country string `bson:"countryOfOrigin" validate:"required"`
}

func (user *User) Validate() error {
	validator := validator.New()
	return validator.Struct(user)
}

// convert user to bson.

// Inserting New User

// Deleting Users

// Updating Users

// Reading Users

// Reading User
