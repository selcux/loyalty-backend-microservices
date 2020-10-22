package consumer

import "go.mongodb.org/mongo-driver/bson/primitive"

type Entity struct {
	ID       primitive.ObjectID `json:"id" bson:"_id" validate:"omitempty"`
	Name     string             `json:"name" bson:"name" validate:"required"`
	Lastname string             `json:"lastname" bson:"lastname" validate:"required"`
	Email    string             `json:"email" bson:"email" validate:"required,email"`
	Wallet   map[string]int     `json:"wallet" bson:"wallet" validate:"omitempty"`
}

type CreateDto struct {
	Name     string `json:"name" bson:"name" validate:"required,min=2"`
	Lastname string `json:"lastname" bson:"lastname" validate:"required,min=2"`
	Email    string `json:"email" bson:"email" validate:"required,email"`
}

type UpdateDto struct {
	Name     string `json:"name" bson:"name" validate:"omitempty,min=2"`
	Lastname string `json:"lastname" bson:"lastname" validate:"omitempty,min=2"`
	Email    string `json:"email" bson:"email" validate:"omitempty,email"`
}

type ItemDto struct {
	ID string `json:"id" bson:"id" validate:"required"`
}
