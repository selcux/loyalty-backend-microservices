package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/di"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/item"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/item/api"
	_ "gitlab.com/adesso-turkey/loyalty-backend-microservices/service/item/docs"
)

// @title Item API
// @description This is the item API of LoyaltyDLT project
// @version 1.0
// @host localhost:80
// @BasePath /
func main() {
	conf := di.InitializeConfig()
	itemService := item.NewItemService()

	go func() {
		err := itemService.Run("", conf.Services["item"].GrpcPort)
		if err != nil {
			log.Fatalf("Unable to serve with gRPC %v", err)
		}
	}()

	srv := server.NewWebServer()
	srv.RegisterRoutes(RegisterRoutes)
	srv.Run("", conf.Services["item"].ApiPort)
}

func RegisterRoutes(e *echo.Echo) {
	itemController := api.NewController()

	v1 := e.Group("/")
	{
		v1.POST("/", itemController.Create)
		v1.GET("/:id", itemController.Read)
		v1.GET("/", itemController.ReadAll)
		v1.PATCH("/:id", itemController.Update)
		v1.DELETE("/:id", itemController.Delete)
	}
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
