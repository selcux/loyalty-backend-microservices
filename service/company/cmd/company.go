package main

import (
	"github.com/labstack/echo"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/di"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/company/api"
)

func main() {
	conf := di.InitializeConfig()
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
