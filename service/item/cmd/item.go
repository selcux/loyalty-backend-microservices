package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
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
	//conf := di.InitializeConfig()
	itemService := item.NewItemService()

	go func() {
		//err := itemService.Run("", conf.Services["item"].GrpcPort)
		err := itemService.Run("", 9104)
		if err != nil {
			log.Fatalf("Unable to serve with gRPC %v", err)
		}
	}()

	srv := server.NewWebServer()
	srv.RegisterRoutes(RegisterRoutes)
	//srv.Run("", conf.Services["item"].ApiPort)
	srv.Run("", 9004)
}

func RegisterRoutes(e *echo.Echo) {
	itemController := api.NewController()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/", itemController.Create)
	e.GET("/:id", itemController.Read)
	e.GET("/", itemController.ReadAll)
	e.PATCH("/:id", itemController.Update)
	e.DELETE("/:id", itemController.Delete)
}
