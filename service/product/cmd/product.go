package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/di"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/product/api"
	_ "gitlab.com/adesso-turkey/loyalty-backend-microservices/service/product/docs"
)

// @title Product API
// @description This is the product API of LoyaltyDLT project
// @version 1.0
// @host localhost:80
// @BasePath /
func main() {
	conf := di.InitializeConfig()
	srv := server.NewWebServer()
	srv.RegisterRoutes(RegisterRoutes)
	srv.Run("", conf.Services["product"].ApiPort)
}

func RegisterRoutes(e *echo.Echo) {
	productController := api.NewController()
	v1 := e.Group("/")
	{
		v1.POST("/", productController.Create)
		v1.GET("/:id", productController.Read)
		v1.GET("/", productController.ReadAll)
		v1.PATCH("/:id", productController.Update)
		v1.DELETE("/:id", productController.Delete)
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)

}
