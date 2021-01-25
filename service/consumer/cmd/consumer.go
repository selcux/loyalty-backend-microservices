package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
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
	//conf := di.InitializeConfig()
	consumerService := consumer.NewConsumerService()

	go func() {
		//err := consumerService.Run("", conf.Services["consumer"].GrpcPort)
		err := consumerService.Run("", 9102)
		if err != nil {
			log.Fatalf("Unable to serve with gRPC %v", err)
		}
	}()

	srv := server.NewWebServer()
	srv.RegisterRoutes(RegisterRoutes)
	//srv.Run("", conf.Services["consumer"].ApiPort)
	srv.Run("", 9002)
}

func RegisterRoutes(e *echo.Echo) {
	consumerController := api.NewController()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/", consumerController.Create)
	e.GET("/:id", consumerController.Read)
	e.GET("/", consumerController.ReadAll)
	e.PATCH("/:id", consumerController.Update)
	e.DELETE("/:id", consumerController.Delete)
	e.PUT("/:id/add", consumerController.Add)
	e.DELETE("/:id/remove", consumerController.Remove)
	//e.POST("/consumers/:id/apply/:campaign_id", consumerController.Apply)
}
