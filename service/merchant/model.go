package merchant

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Merchant struct {
	ID       primitive.ObjectID `json:"id" bson:"_id" validate:"omitempty"`
	Name     string             `json:"name" bson:"name" validate:"required"`
	Location Location           `json:"location" bson:"location" validate:"required"`
}

type Location struct {
	Latitude  float64 `json:"latitude" bson:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" bson:"longitude" validate:"required"`
}

type CreateDto struct {
	Name     string   `json:"name" bson:"name" validate:"required,min=2"`
	Location Location `json:"location" bson:"location" validate:"required"`
}

type UpdateDto struct {
	Name string `json:"name" bson:"name" validate:"required,min=2"`
	// TODO: Location didn't change
	Location Location `json:"location" bson:"location" validate:"required"`
}
