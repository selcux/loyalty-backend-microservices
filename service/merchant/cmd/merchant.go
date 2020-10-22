package main

import (
	"github.com/labstack/echo"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/config"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/merchant/api"
	"log"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Unable to read loyalty.yaml %v", err)
	}

	srv := server.NewWebServer()
	srv.RegisterRoutes(RegisterRoutes)
	srv.Run("", conf.Services["merchant"].ApiPort)
}

func RegisterRoutes(e *echo.Echo) {
	merchantController := api.NewController()
	e.POST("/", merchantController.Create)
	e.GET("/:id", merchantController.Read)
	e.GET("/", merchantController.ReadAll)
	e.PUT("/:id", merchantController.Update)
	e.DELETE("/:id", merchantController.Delete)
}
