package main

import (
	"github.com/labstack/echo"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/config"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/company/api"
	"log"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Unable to read loyalty.yaml %v", err)
	}

	srv := server.NewWebServer()
	srv.RegisterRoutes(RegisterRoutes)
	srv.Run("", conf.Services["company"].ApiPort)
}

func RegisterRoutes(e *echo.Echo) {
	companyController := api.NewController()
	e.POST("/", companyController.Create)
	e.GET("/:id", companyController.Read)
	e.GET("/", companyController.ReadAll)
	e.PUT("/:id", companyController.Update)
	e.DELETE("/:id", companyController.Delete)
}
