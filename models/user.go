package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `bson:"email" validate:"required,email"`
	Password string             `bson:"password" validate:"required"`
	Countries []string           `bson:"countries"`
}

func ValidateUser(user *User) error {
	validate := validator.New()
	return validate.Struct(user)
}
