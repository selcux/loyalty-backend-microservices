package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/util"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/company"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

// Create godoc
// @Summary Create a company
// @Description Create a new company
// @Tags company
// @Accept json
// @Produce json
// @Param company body company.CreateDto true "New Company"
// @Success 201 {object} company.Company
// @Failure 400 {object} util.ErrorResponse
// @Router /companies [post]
func (c *Controller) Create(ctx echo.Context) error {
	vm := new(company.CreateDto)
	if err := ctx.Bind(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	if err := ctx.Validate(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	cdb, err := company.NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := cdb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	companyData, err := cdb.Create(vm)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusCreated, util.CreateOkResponse(companyData))
}

// Read godoc
// @Summary Read a company data
// @Description Get a company data
// @Tags company
// @Produce json
// @Param id path string true "Company ID"
// @Success 200 {object} company.Company
// @Failure 400 {object} util.ErrorResponse
// @Router /companies/{id} [get]
func (c *Controller) Read(ctx echo.Context) error {
	cdb, err := company.NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := cdb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	paramId := ctx.Param("id")
	companyData, err := cdb.Read(paramId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, util.CreateOkResponse(companyData))
}

// ReadAll godoc
// @Summary Read all company data
// @Description Get all companies
// @Tags company
// @Produce json
// @Success 200 {object} []company.Company
// @Failure 400 {object} util.ErrorResponse
// @Router /companies [get]
func (c *Controller) ReadAll(ctx echo.Context) error {
	cdb, err := company.NewDb()
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

// Update godoc
// @Summary Update a company data
// @Description Update a company data
// @Tags company
// @Accept json
// @Produce json
// @Param id path string true "Company ID"
// @Param company body company.UpdateDto true "Update Company"
// @Success 200 {object} company.Company
// @Failure 400 {object} util.ErrorResponse
// @Router /companies/{id} [put]
func (c *Controller) Update(ctx echo.Context) error {
	vm := new(company.UpdateDto)
	if err := ctx.Bind(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	if err := ctx.Validate(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	cdb, err := company.NewDb()
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

// Delete godoc
// @Summary Delete a company data
// @Description Delete a company data
// @Tags company
// @Produce json
// @Param id path string true "Company ID"
// @Success 204
// @Failure 400 {object} util.ErrorResponse
// @Router /companies/{id} [delete]
func (c *Controller) Delete(ctx echo.Context) error {
	cdb, err := company.NewDb()
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
