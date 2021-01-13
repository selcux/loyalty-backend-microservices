package main

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/di"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/campaign/api"
)

func main() {
	conf := di.InitializeConfig()
	srv := server.NewWebServer()
	srv.RegisterRoutes(RegisterRoutes)
	srv.Run("", conf.Services["campaign"].ApiPort)
}

func RegisterRoutes(e *echo.Echo) {
	campaignController := api.NewController()
	e.POST("/", campaignController.Create)
}
