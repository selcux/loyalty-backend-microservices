package util

import (
	"encoding/json"
	"github.com/labstack/echo"
)

type CrudControllerInterface interface {
	Create(ctx echo.Context) error
	Read(ctx echo.Context) error
	ReadAll(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type OkResponse struct {
	Data interface{} `json:"data" bson:"data"`
}

type ErrorResponse struct {
	Message string `json:"message" bson:"message"`
}

func CreateOkResponse(data interface{}) *OkResponse {
	return &OkResponse{Data: data}
}

func CreateErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		Message: err.Error(),
	}
}

type OkResponseJson struct {
	Data json.RawMessage `json:"data" bson:"data"`
}
