package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/product/api"
	_ "gitlab.com/adesso-turkey/loyalty-backend-microservices/service/product/docs"
)

// @title Product API
// @description This is the product API of LoyaltyDLT project
// @version 0.1
// @BasePath /
func main() {
	//conf := di.InitializeConfig()
	srv := server.NewWebServer()
	srv.RegisterRoutes(RegisterRoutes)
	//srv.Run("", conf.Services["product"].ApiPort)
	srv.Run("", 9003)
}

func RegisterRoutes(e *echo.Echo) {
	productController := api.NewController()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/", productController.Create)
	e.GET("/:id", productController.Read)
	e.GET("/", productController.ReadAll)
	e.PATCH("/:id", productController.Update)
	e.DELETE("/:id", productController.Delete)
}
