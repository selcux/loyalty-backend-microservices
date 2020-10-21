package company

import "go.mongodb.org/mongo-driver/bson/primitive"

type Company struct {
	ID   primitive.ObjectID `json:"id" bson:"_id" validate:"omitempty"`
	Name string             `json:"name" bson:"name" validate:"required"`
}

type CreateDto struct {
	Name string `json:"name" bson:"name" validate:"required,min=3"`
}

type UpdateDto struct {
	Name string `json:"name" bson:"name" validate:"required,min=3"`
}
