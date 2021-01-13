package api

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/util"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/campaign/model"
	"net/http"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Create(ctx echo.Context) error {
	vm := new(model.Campaign)
	if err := ctx.Bind(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	if err := ctx.Validate(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusCreated, util.CreateOkResponse(vm))
}

func (c *Controller) Read(ctx echo.Context) error {
	panic("implement me")
}

func (c *Controller) ReadAll(ctx echo.Context) error {
	panic("implement me")
}

/*
func (c *Controller) Update(ctx echo.Context) error {
	panic("implement me")
}

func (c *Controller) Delete(ctx echo.Context) error {
	panic("implement me")
}
*/
