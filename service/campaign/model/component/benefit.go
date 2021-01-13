package component

import "gitlab.com/adesso-turkey/loyalty-backend-microservices/service/campaign/model"

type Benefit struct {
	model.Component
	Benefits map[string]int `json:"benefits" validate:"required"`
}
