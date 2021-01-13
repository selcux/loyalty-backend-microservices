package component

import "gitlab.com/adesso-turkey/loyalty-backend-microservices/service/campaign/model"

type Point struct {
	model.Component
	Items map[string]int `json:"items" validate:"required"`
}
