package main

import (
	"fmt"
	"github.com/labstack/echo"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/util"
	"net/http"
)

func main() {
	srv := server.NewWebServer()
	srv.RegisterRoutes(RegisterRoutes)
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
