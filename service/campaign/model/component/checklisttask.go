package component

import "gitlab.com/adesso-turkey/loyalty-backend-microservices/service/campaign/model"

type ChecklistTask struct {
	model.Component
	Tasks map[string]int `json:"tasks" validate:"required"`
}
