package main

import (
	"github.com/labstack/echo"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/config"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/product/api"
	"log"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Unable to read loyalty.yaml %v", err)
	}

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
