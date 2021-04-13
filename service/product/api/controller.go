package api

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/client"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/util"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/grpc/item"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/product"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

// Create godoc
// @Summary Create a product
// @Description Create a new product
// @Tags product
// @Accept json
// @Produce json
// @Param product body product.CreateDto true "New Product"
// @Success 201 {object} product.Product
// @Failure 400 {object} util.ErrorResponse
// @Router /products [post]
func (c *Controller) Create(ctx echo.Context) error {
	vm := new(product.CreateDto)
	vm.Company = ctx.Request().Header.Get("Company")

	if err := ctx.Bind(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	if err := ctx.Validate(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	pdb, err := product.NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := pdb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	productData, err := pdb.Create(vm)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	itemClient := client.NewItemService()
	err = itemClient.Connect()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer itemClient.Close()

	err = itemClient.Create(&item.CreateItemDto{
		Name:    productData.Name,
		Company: productData.Company.Hex(),
		Product: productData.ID.Hex(),
		Point:   int32(productData.Point),
		Code:    productData.Code,
	})

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusCreated, util.CreateOkResponse(productData))
}

// Read godoc
// @Summary Read a product data
// @Description Get a product data
// @Tags product
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} product.Product
// @Failure 400 {object} util.ErrorResponse
// @Router /products/{id} [get]
func (c *Controller) Read(ctx echo.Context) error {
	pdb, err := product.NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := pdb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	paramId := ctx.Param("id")
	product, err := pdb.Read(paramId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, util.CreateOkResponse(product))
}

// ReadAll godoc
// @Summary Read all product data
// @Description Get all products
// @Tags product
// @Produce json
// @Success 200 {object} []product.Product
// @Failure 400 {object} util.ErrorResponse
// @Router /products [get]
func (c *Controller) ReadAll(ctx echo.Context) error {
	log.Println("Got read request")
	pdb, err := product.NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}
	log.Println("Opened DB connection...")
	defer func() {
		if err := pdb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	products, err := pdb.ReadAll()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}
	log.Println("Read elements of DB")

	return ctx.JSON(http.StatusOK, util.CreateOkResponse(products))
}

// Update godoc
// @Summary Update a product data
// @Description Update a product data
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param product body product.UpdateDto true "Update Product"
// @Success 200 {object} product.Product
// @Failure 400 {object} util.ErrorResponse
// @Router /products/{id} [patch]
func (c *Controller) Update(ctx echo.Context) error {
	vm := new(product.UpdateDto)
	if err := ctx.Bind(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	if err := ctx.Validate(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	pdb, err := product.NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := pdb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	paramId := ctx.Param("id")

	product, err := pdb.Update(paramId, vm)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, util.CreateOkResponse(product))
}

// Delete godoc
// @Summary Delete a product data
// @Description Delete a product data
// @Tags product
// @Produce json
// @Param id path string true "Product ID"
// @Success 204
// @Failure 400 {object} util.ErrorResponse
// @Router /products/{id} [delete]
func (c *Controller) Delete(ctx echo.Context) error {
	pdb, err := product.NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := pdb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	paramId := ctx.Param("id")
	err = pdb.Delete(paramId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
