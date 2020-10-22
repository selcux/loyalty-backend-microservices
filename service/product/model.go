package product

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID      primitive.ObjectID `json:"id" bson:"_id" validate:"omitempty"`
	Company primitive.ObjectID `json:"company" bson:"company,required" validate:"required"`
	Name    string             `json:"name" bson:"name" validate:"required"`
	Point   int                `json:"point" bson:"point,required" validate:"required"`
	Code    string             `json:"code" bson:"code" validate:"required"`
}

type CreateDto struct {
	Name string `json:"name" bson:"name" validate:"required,min=2"`
	//TODO: string?
	Company string `json:"company" bson:"company" validate:"required"`
	Point   int    `json:"point" bson:"point" validate:"required"`
	Code    string `json:"code" bson:"code" validate:"required"`
}

type UpdateDto struct {
	Name string `json:"name" bson:"name" validate:"omitempty,min=2"`
	Code string `json:"code" bson:"code" validate:"required"`
}

type CreateProduct struct {
	Name    string             `json:"name" bson:"name" validate:"required,min=2"`
	Company primitive.ObjectID `json:"company" bson:"company,required" validate:"required"`
	Point   int                `json:"point" bson:"point" validate:"required"`
	Code    string             `json:"code" bson:"code" validate:"required"`
}
