package main

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/di"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/merchant/api"
)

func main() {
	conf := di.InitializeConfig()
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
