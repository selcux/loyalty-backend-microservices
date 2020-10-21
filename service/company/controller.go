package company

import (
	"github.com/labstack/echo"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/util"
	"net/http"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Create(ctx echo.Context) error {
	vm := new(CreateDto)
	if err := ctx.Bind(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	if err := ctx.Validate(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	cdb, err := NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := cdb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	company, err := cdb.Create(vm)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusCreated, util.CreateOkResponse(company))
}

func (c *Controller) Read(ctx echo.Context) error {
	cdb, err := NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := cdb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	paramId := ctx.Param("id")
	company, err := cdb.Read(paramId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, util.CreateOkResponse(company))
}

func (c *Controller) ReadAll(ctx echo.Context) error {
	cdb, err := NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := cdb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	companies, err := cdb.ReadAll()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, util.CreateOkResponse(companies))
}

func (c *Controller) Update(ctx echo.Context) error {
	vm := new(UpdateDto)
	if err := ctx.Bind(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	if err := ctx.Validate(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	cdb, err := NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := cdb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	paramId := ctx.Param("id")

	err = cdb.Update(paramId, vm)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, util.CreateOkResponse(nil))
}

func (c *Controller) Delete(ctx echo.Context) error {
	cdb, err := NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := cdb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	paramId := ctx.Param("id")
	err = cdb.Delete(paramId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
