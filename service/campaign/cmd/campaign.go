package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/di"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/campaign/api"
	_ "gitlab.com/adesso-turkey/loyalty-backend-microservices/service/campaign/docs"
)

// @title Campaign API
// @description This is the campaign API of LoyaltyDLT project
// @version 1.0
// @host localhost:80
// @BasePath /
func main() {
	conf := di.InitializeConfig()
	srv := server.NewWebServer()
	srv.RegisterRoutes(RegisterRoutes)
	srv.Run("", conf.Services["campaign"].ApiPort)
}

func RegisterRoutes(e *echo.Echo) {
	campaignController := api.NewController()
	v1 := e.Group("/")
	{
		v1.POST("/", campaignController.Create)
	}

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
