package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/util"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/item"
)

type Controller struct {
}

// Create godoc
// @Summary Create an item data
// @Description Create an item data
// @Tags item
// @Accept json
// @Produce json
// @Param item body item.CreateDto true "Create an item"
// @Success 201 {object} item.Entity
// @Failure 400 {object} util.ErrorResponse
// @Router /items [post]
func (c *Controller) Create(ctx echo.Context) error {
	vm := new(item.CreateDto)
	if err := ctx.Bind(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	if err := ctx.Validate(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	idb, err := item.NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := idb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	itm, err := idb.Create(vm)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusCreated, util.CreateOkResponse(itm))
}

// Read godoc
// @Summary Read an item data
// @Description Read an item data
// @Tags item
// @Produce json
// @Param id path string true "Item ID"
// @Success 200 {object} item.Entity
// @Failure 400 {object} util.ErrorResponse
// @Router /items/{id} [get]
func (c *Controller) Read(ctx echo.Context) error {
	idb, err := item.NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := idb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	itemId := ctx.Param("id")
	itm, err := idb.Read(itemId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, util.CreateOkResponse(itm))
}

// ReadAll godoc
// @Summary Read all items
// @Description Read all items
// @Tags item
// @Produce json
// @Success 200 {object} []item.Entity
// @Failure 400 {object} util.ErrorResponse
// @Router /items [get]
func (c *Controller) ReadAll(ctx echo.Context) error {
	idb, err := item.NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := idb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	items, err := idb.ReadAll()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, util.CreateOkResponse(items))
}

// Update godoc
// @Summary Update an item
// @Description Update an item
// @Tags item
// @Accept json
// @Produce json
// @Param id path string true "Item ID"
// @Param item body item.UpdateDto true "Update all items"
// @Success 200 {object} item.Entity
// @Failure 400 {object} util.ErrorResponse
// @Router /items/{id} [patch]
func (c *Controller) Update(ctx echo.Context) error {
	vm := new(item.UpdateDto)
	if err := ctx.Bind(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	if err := ctx.Validate(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	idb, err := item.NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := idb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	paramId := ctx.Param("id")

	itm, err := idb.Update(paramId, vm)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, util.CreateOkResponse(itm))
}

// Delete godoc
// @Summary Delete an item
// @Description Delete an item
// @Tags item
// @Produce json
// @Param id path string true "Item ID"
// @Success 204
// @Failure 400 {object} util.ErrorResponse
// @Router /items/{id} [delete]
func (c *Controller) Delete(ctx echo.Context) error {
	idb, err := item.NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := idb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	paramId := ctx.Param("id")
	err = idb.Delete(paramId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusNoContent, nil)
}

func NewController() *Controller {
	return &Controller{}
}
