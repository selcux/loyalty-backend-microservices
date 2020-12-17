package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/di"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/consumer"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/consumer/api"
)

func main() {
	conf := di.InitializeConfig()
	consumerService := consumer.NewConsumerService()

	go func() {
		err := consumerService.Run("", conf.Services["consumer"].GrpcPort)
		if err != nil {
			log.Fatalf("Unable to serve with gRPC %v", err)
		}
	}()

	srv := server.NewWebServer()
	srv.RegisterRoutes(RegisterRoutes)
	srv.Run("", conf.Services["consumer"].ApiPort)
}

func RegisterRoutes(e *echo.Echo) {
	consumerController := api.NewController()
	e.POST("/", consumerController.Create)
	e.GET("/:id", consumerController.Read)
	e.GET("/", consumerController.ReadAll)
	e.PATCH("/:id", consumerController.Update)
	e.DELETE("/:id", consumerController.Delete)
	e.PUT("/:id/add", consumerController.Add)
	e.DELETE("/:id/remove", consumerController.Remove)
	//e.POST("/consumers/:id/apply/:campaign_id", consumerController.Apply)
}
