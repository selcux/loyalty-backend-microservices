package api

import (
	"net/http"

	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/util"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/merchant"

	"github.com/labstack/echo/v4"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

// Create godoc
// @Summary Create a merchant
// @Description Create a new merchant
// @Tags merchant
// @Accept json
// @Produce json
// @Param merchant body merchant.CreateDto true "New Merchant"
// @Success 201 {object} merchant.Merchant
// @Failure 400 {object} util.ErrorResponse
// @Router /merchants [post]
func (c *Controller) Create(ctx echo.Context) error {
	vm := new(merchant.CreateDto)

	if err := ctx.Bind(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	if err := ctx.Validate(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	mdb, err := merchant.NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := mdb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	merchantData, err := mdb.Create(vm)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusCreated, util.CreateOkResponse(merchantData))
}

// Read godoc
// @Summary Read a merchant data
// @Description Get a merchant data
// @Tags merchant
// @Produce json
// @Param id path string true "Merchant ID"
// @Success 200 {object} merchant.Merchant
// @Failure 400 {object} util.ErrorResponse
// @Router /merchants/{id} [get]
func (c *Controller) Read(ctx echo.Context) error {
	mdb, err := merchant.NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := mdb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	merchantId := ctx.Param("id")
	merchantData, err := mdb.Read(merchantId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, util.CreateOkResponse(merchantData))
}

// ReadAll godoc
// @Summary Read all merchant data
// @Description Get all merchants
// @Tags merchant
// @Produce json
// @Success 200 {object} []merchant.Merchant
// @Failure 400 {object} util.ErrorResponse
// @Router /merchants [get]
func (c *Controller) ReadAll(ctx echo.Context) error {
	mdb, err := merchant.NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := mdb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	merchants, err := mdb.ReadAll()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, util.CreateOkResponse(merchants))
}

// Update godoc
// @Summary Update a merchant data
// @Description Update a merchant data
// @Tags merchant
// @Accept json
// @Produce json
// @Param merchant body merchant.UpdateDto true "Update Merchant"
// @Param id path string true "Merchant ID"
// @Success 200 {object} merchant.Merchant
// @Failure 400 {object} util.ErrorResponse
// @Router /merchants/{id} [put]
func (c *Controller) Update(ctx echo.Context) error {
	vm := new(merchant.UpdateDto)
	if err := ctx.Bind(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	if err := ctx.Validate(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	mdb, err := merchant.NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := mdb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	paramId := ctx.Param("id")

	merchantData, err := mdb.Update(paramId, vm)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}
	return ctx.JSON(http.StatusOK, util.CreateOkResponse(merchantData))
}

// Delete godoc
// @Summary Delete a merchant data
// @Description Delete a merchant data
// @Tags merchant
// @Produce json
// @Param id path string true "Merchant ID"
// @Success 204
// @Failure 400 {object} util.ErrorResponse
// @Router /merchants/{id} [delete]
func (c *Controller) Delete(ctx echo.Context) error {
	mdb, err := merchant.NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := mdb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()
	paramId := ctx.Param("id")
	err = mdb.Delete(paramId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
