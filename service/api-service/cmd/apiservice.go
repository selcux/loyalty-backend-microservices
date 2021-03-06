package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/util"
	"net/http"
)

func main() {
	//conf := di.InitializeConfig()
	srv := server.NewWebServer()
	srv.RegisterRoutes(RegisterRoutes)
	//srv.Run("", conf.Services["api_service"].ApiPort)
	srv.Run("", 9000)
}
func RegisterRoutes(e *echo.Echo) {
	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, util.CreateOkResponse(struct {
			Message string `json:"message"`
		}{
			Message: "It' alive!",
		}))
	})

	e.GET("/:name", func(ctx echo.Context) error {
		name := ctx.Param("name")
		return ctx.JSON(http.StatusOK, util.CreateOkResponse(struct {
			Message string `json:"message"`
		}{
			Message: fmt.Sprintf("Hello %s!", name),
		}))
	})
}
