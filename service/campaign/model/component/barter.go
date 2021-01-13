package component

import "gitlab.com/adesso-turkey/loyalty-backend-microservices/service/campaign/model"

type Barter struct {
	model.Component
	ChecklistTask
	Benefit
	Expiration bool `json:"expiration"` // Not functional yet
}
