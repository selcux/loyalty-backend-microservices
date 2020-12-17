package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/di"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/item"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/item/api"
)

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
	e.POST("/", itemController.Create)
	e.GET("/:id", itemController.Read)
	e.GET("/", itemController.ReadAll)
	e.PATCH("/:id", itemController.Update)
	e.DELETE("/:id", itemController.Delete)
}
