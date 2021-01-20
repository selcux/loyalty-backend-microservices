package api

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/client"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/util"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/consumer"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

// Create godoc
// @Summary Create a consumer data
// @Description Create a consumer data
// @Tags consumer
// @Accept json
// @Produce json
// @Param consumer body consumer.CreateDto true "Create consumer"
// @Success 201 {object} consumer.CreateDto
// @Failure 400 {object} HTTPError
// @Router /consumer [post]
func (c *Controller) Create(ctx echo.Context) error {
	vm := new(consumer.CreateDto)
	if err := ctx.Bind(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	if err := ctx.Validate(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	cdb, err := consumer.NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := cdb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	consumerData, err := cdb.Create(vm)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusCreated, util.CreateOkResponse(consumerData))
}

// Read godoc
// @Summary Read a consumer data
// @Description Get a consumer data
// @Tags consumer
// @Accept json
// @Produce json
// @Param consumer body model.Consumer true "Read consumer"
// @Success 201 {object} model.Consumer
// @Failure 400 {object} HTTPError
// @Router /consumer [get]
func (c *Controller) Read(ctx echo.Context) error {
	cdb, err := consumer.NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := cdb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	paramId := ctx.Param("id")
	consumerData, err := cdb.Read(paramId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, util.CreateOkResponse(consumerData))
}

// ReadAll godoc
// @Summary Read all consumer data
// @Description Get all consumer data
// @Tags consumer
// @Accept json
// @Produce json
// @Param consumer body []model.Consumer true "Read all consumer"
// @Success 201 {object} []model.Consumer
// @Failure 400 {object} HTTPError
// @Router /consumer [get]
func (c *Controller) ReadAll(ctx echo.Context) error {
	cdb, err := consumer.NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := cdb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	consumers, err := cdb.ReadAll()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, util.CreateOkResponse(consumers))
}

// Update godoc
// @Summary Update a consumer data
// @Description Update a consumer data
// @Tags consumer
// @Accept json
// @Produce json
// @Param consumer body consumer.UpdateDto true "Update a consumer"
// @Success 201 {object} consumer.UpdateDto
// @Failure 400 {object} HTTPError
// @Router /consumer [patch]
func (c *Controller) Update(ctx echo.Context) error {
	vm := new(consumer.UpdateDto)
	if err := ctx.Bind(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	if err := ctx.Validate(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	cdb, err := consumer.NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := cdb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	paramId := ctx.Param("id")

	consumerData, err := cdb.Update(paramId, vm)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, util.CreateOkResponse(consumerData))
}

// Delete godoc
// @Summary Delete a consumer data
// @Description Delete a consumer data
// @Tags consumer
// @Accept json
// @Produce json
// @Param consumer body model.Consumer true "Delete a consumer"
// @Success 201 {object} model.Consumer
// @Failure 400 {object} HTTPError
// @Router /consumer [delete]
func (c *Controller) Delete(ctx echo.Context) error {
	cdb, err := consumer.NewDb()
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

// Add godoc
// @Summary Add an item to consumer's wallet
// @Description Update a consumer data by adding an item to consumer's wallet
// @Tags consumer
// @Accept json
// @Produce json
// @Param consumer body consumer.ItemDto true "Update a consumer wallet"
// @Success 201 {object} consumer.ItemDto
// @Failure 400 {object} HTTPError
// @Router /consumer [put]
func (c *Controller) Add(ctx echo.Context) error {
	vm := new(consumer.ItemDto)
	if err := ctx.Bind(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	if err := ctx.Validate(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	itemClient := client.NewItemService()
	err := itemClient.Connect()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer itemClient.Close()

	found, err := itemClient.ItemExists(vm.ID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}
	if !found {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(errors.New("invalid item")))
	}

	cdb, err := consumer.NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := cdb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	paramId := ctx.Param("id")
	err = cdb.AddToWallet(paramId, vm.ID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusNoContent, nil)
}

// Remove godoc
// @Summary Remove an item from consumer's wallet
// @Description Update a consumer data by remove an item from consumer's wallet
// @Tags consumer
// @Accept json
// @Produce json
// @Param consumer body consumer.ItemDto true "Remove from the consumer wallet"
// @Success 201 {object} consumer.ItemDto
// @Failure 400 {object} HTTPError
// @Router /consumer [delete]
func (c *Controller) Remove(ctx echo.Context) error {
	vm := new(consumer.ItemDto)
	if err := ctx.Bind(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	if err := ctx.Validate(vm); err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	itemClient := client.NewItemService()
	err := itemClient.Connect()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer itemClient.Close()

	found, err := itemClient.ItemExists(vm.ID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}
	if !found {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(errors.New("invalid item")))
	}

	cdb, err := consumer.NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := cdb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	paramId := ctx.Param("id")
	err = cdb.RemoveFromWallet(paramId, vm.ID, 1)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusNoContent, nil)
}

/* func (c *Controller) Apply(ctx echo.Context) error {
	cdb, err := consumerDb.NewDb()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	defer func() {
		if err := cdb.Close(); err != nil {
			ctx.Logger().Error(err)
		}
	}()

	paramId := ctx.Param("id")
	campaignId := ctx.Param("campaign_id")

	comm := util.NewServiceComm()
	campaignData, err := retrieveCampaign(comm, campaignId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	consumer, err := cdb.Read(paramId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	for _, cc := range campaignData.Components {
		point := new(component.Point)
		barter := new(component.Barter)

		comp, err := util.ConvertToICampaignComponent(cc, point, barter)
		if err != nil {
			log.Println(err)
			return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
			//continue
		}

		decrementItems := comp.ToRemove(consumer.Wallet)

		for k, v := range decrementItems {
			werr := cdb.RemoveFromWallet(paramId, k, v)
			if werr != nil {
				return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
			}
		}
	}

	consumer, err = cdb.Read(paramId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, util.CreateErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, util.CreateOkResponse(consumer))
}

func retrieveCampaign(comm util.IServiceComm, campaignId string) (*campaign.Campaign, error) {
	result, err := comm.Get(util.CampaignPath, campaignId)
	if err != nil {
		return nil, err
	}

	resultStr := string(result)
	response := new(util.OkResponseJson)
	err = json.Unmarshal([]byte(resultStr), response)
	if err != nil {
		return nil, err
	}

	campaignData := new(campaign.Campaign)
	err = json.Unmarshal(response.Data, campaignData)
	//log.Println("Campaign Data", fmt.Sprintf("%+v", campaignData))
	return campaignData, err
}
*/
