package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/di"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/company/api"
	_ "gitlab.com/adesso-turkey/loyalty-backend-microservices/service/company/docs"
)

// @title Company API
// @description This is the company API of LoyaltyDLT project
// @version 1.0
// @host localhost:80
// @BasePath /
func main() {
	conf := di.InitializeConfig()
	srv := server.NewWebServer()
	srv.RegisterRoutes(RegisterRoutes)
	srv.Run("", conf.Services["company"].ApiPort)
}

func RegisterRoutes(e *echo.Echo) {
	companyController := api.NewController()

	v1 := e.Group("/")
	{
		v1.POST("/", companyController.Create)
		v1.GET("/:id", companyController.Read)
		v1.GET("/", companyController.ReadAll)
		v1.PUT("/:id", companyController.Update)
		v1.DELETE("/:id", companyController.Delete)
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
