package api

import (
	"errors"
	"github.com/labstack/echo"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/client"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/util"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/consumer"
	"net/http"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

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
