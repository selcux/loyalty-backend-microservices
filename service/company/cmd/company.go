package main

import (
	"github.com/labstack/echo"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/company"
)

func main() {
	srv := server.NewWebServer()
	srv.RegisterRoutes(RegisterRoutes)
	srv.Run("", 9001)
}

func RegisterRoutes(e *echo.Echo) {
	companyController := company.NewController()
	e.POST("/", companyController.Create)
	e.GET("/:id", companyController.Read)
	e.GET("/", companyController.ReadAll)
	e.PUT("/:id", companyController.Update)
	e.DELETE("/:id", companyController.Delete)
}
