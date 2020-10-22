package item

import "go.mongodb.org/mongo-driver/bson/primitive"

type Entity struct {
	ID      primitive.ObjectID `json:"id" bson:"_id" validate:"omitempty"`
	Name    string             `json:"name" bson:"name" validate:"required"`
	Company primitive.ObjectID `json:"company" bson:"company,required" validate:"required"`
	Product primitive.ObjectID `json:"product" bson:"product,required" validate:"required"`
	Group   string             `json:"group" bson:"group"`
	Point   int                `json:"point" bson:"point,required" validate:"required"`
	Code    string             `json:"code" bson:"code" validate:"required"`
}

type CreateDto struct {
	Name    string             `json:"name" bson:"name" validate:"required,min=2"`
	Company primitive.ObjectID `json:"company" bson:"company" validate:"omitempty"`
	Product primitive.ObjectID `json:"product" bson:"product" validate:"omitempty"`
	Point   int                `json:"point" bson:"point" validate:"required"`
	Code    string             `json:"code" bson:"code" validate:"omitempty"`
}

type UpdateDto struct {
	Name  string `json:"name" bson:"name" validate:"required,min=2"`
	Point int    `json:"point" bson:"point" validate:"required"`
}
