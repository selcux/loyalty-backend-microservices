package main

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/company/api"
	_ "gitlab.com/adesso-turkey/loyalty-backend-microservices/service/company/docs"
)

// @title Company API
// @description This is the company API of LoyaltyDLT project
// @version 0.1
// @BasePath /
func main() {
	//conf := di.InitializeConfig()
	srv := server.NewWebServer()
	srv.RegisterRoutes(RegisterRoutes)
	srv.Run("", 9001)
}

func RegisterRoutes(e *echo.Echo) {
	companyController := api.NewController()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/", companyController.Create)
	e.GET("/:id", companyController.Read)
	e.GET("/", companyController.ReadAll)
	e.PUT("/:id", companyController.Update)
	e.DELETE("/:id", companyController.Delete)
}
