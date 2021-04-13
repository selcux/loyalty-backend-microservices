package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/merchant/api"
	_ "gitlab.com/adesso-turkey/loyalty-backend-microservices/service/merchant/docs"
)

// @title Merchant API
// @description This is the merchant API of LoyaltyDLT project
// @version 1.0
// @BasePath /
func main() {
	//conf := di.InitializeConfig()
	srv := server.NewWebServer()
	srv.RegisterRoutes(RegisterRoutes)
	//srv.Run("", conf.Services["merchant"].ApiPort)
	srv.Run("", 9006)
}

func RegisterRoutes(e *echo.Echo) {
	merchantController := api.NewController()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/", merchantController.Create)
	e.GET("/:id", merchantController.Read)
	e.GET("/", merchantController.ReadAll)
	e.PUT("/:id", merchantController.Update)
	e.DELETE("/:id", merchantController.Delete)
}
