package api

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/campaign/model"
	"net/http"
)

type handler struct {
	campaign *model.Campaign
}

func (h *handler) createCampaign(c echo.Context) error {
	camp := new(model.Campaign)
	if err := c.Bind(camp); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, camp)
}
