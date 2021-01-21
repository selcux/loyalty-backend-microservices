package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/di"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/consumer"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/consumer/api"
	_ "gitlab.com/adesso-turkey/loyalty-backend-microservices/service/consumer/docs"
)

// @title Consumer API
// @description This is the consumer API of LoyaltyDLT project
// @version 1.0
// @host localhost:80
// @BasePath /
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
	v1 := e.Group("/")
	{
		v1.POST("/", consumerController.Create)
		v1.GET("/:id", consumerController.Read)
		v1.GET("/", consumerController.ReadAll)
		v1.PATCH("/:id", consumerController.Update)
		v1.DELETE("/:id", consumerController.Delete)
		v1.PUT("/:id/add", consumerController.Add)
		v1.DELETE("/:id/remove", consumerController.Remove)
	}
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	//e.POST("/consumers/:id/apply/:campaign_id", consumerController.Apply)
}
