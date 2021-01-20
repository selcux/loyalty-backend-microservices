package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/di"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/merchant/api"
)

// @title Merchant API
// @description This is the merchant API of LoyaltyDLT project
// @version 1.0
// @host localhost:80
// @BasePath /api/v1
func main() {
	conf := di.InitializeConfig()
	srv := server.NewWebServer()
	srv.RegisterRoutes(RegisterRoutes)
	srv.Run("", conf.Services["merchant"].ApiPort)
}

func RegisterRoutes(e *echo.Echo) {
	merchantController := api.NewController()
	v1 := e.Group("/api/v1")
	{
		v1.POST("/", merchantController.Create)
		v1.GET("/:id", merchantController.Read)
		v1.GET("/", merchantController.ReadAll)
		v1.PUT("/:id", merchantController.Update)
		v1.DELETE("/:id", merchantController.Delete)
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
