package main

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/di"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/product/api"
)

func main() {
	conf := di.InitializeConfig()
	srv := server.NewWebServer()
	srv.RegisterRoutes(RegisterRoutes)
	srv.Run("", conf.Services["product"].ApiPort)
}

func RegisterRoutes(e *echo.Echo) {
	productController := api.NewController()
	e.POST("/", productController.Create)
	e.GET("/:id", productController.Read)
	e.GET("/", productController.ReadAll)
	e.PATCH("/:id", productController.Update)
	e.DELETE("/:id", productController.Delete)
}
