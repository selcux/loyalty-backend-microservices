package component

import "gitlab.com/adesso-turkey/loyalty-backend-microservices/service/campaign/model"

type TierBenefit struct {
	model.Component
	Paid  bool           `json:"paid"`
	Items map[string]int `json:"items" validate:"required"`
}
