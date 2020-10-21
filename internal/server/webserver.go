package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

type WebServer struct {
	e *echo.Echo
}

func NewWebServer() *WebServer {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	e.Validator = &CustomValidator{validator: validator.New()}

	return &WebServer{
		e: e,
	}
}

func (server *WebServer) RegisterRoutes(f func(e *echo.Echo)) {
	f(server.e)
}

func (server *WebServer) Run(host string, port uint) {
	// Start server
	go func() {
		if err := server.e.Start(fmt.Sprintf("%s:%d", host, port)); err != nil {
			server.e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.e.Shutdown(ctx); err != nil {
		server.e.Logger.Fatal(err)
	}
}
