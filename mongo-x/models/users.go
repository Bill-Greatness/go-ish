package models

import (
	"github.com/go-playground/validator/v10"
)

// user model definitions.
type User struct {
	Name    string `json:"name" validate:"required"`
	Age     int32  `json:"age" validate:"required"`
	Email   string `json:"email"`
	Country string `json:"countryOfOrigin" validate:"required"`
}

func (user *User) Validate() error {
	validator := validator.New()
	return validator.Struct(user)
}

// Inserting New User

// Deleting Users

// Updating Users

// Reading Users

// Reading User
