package model

type Type string

const (
	Point             Type = "point"
	Barter            Type = "barter"
	Distribution      Type = "distribution"
	Lottery           Type = "lottery"
	AdditionalBenefit Type = "additional_benefit"
	Enrollment        Type = "enrollment"
	Ordered           Type = "ordered"
	Task              Type = "task"
	Benefit           Type = "benefit"
	ChecklistLocation Type = "location"
)

type Container struct {
	ComponentType Type `json:"component_type" validate:"required"`
	Order         int  `json:"order" validate:"required"`
	//Component     json.RawMessage `json:"component" validate:"required"`
	Component ComponentComponent `json:"component" validate:"required"`
}

type Campaign struct {
	Name         string      `json:"name" validate:"required"`
	PublicKey    string      `json:"public_key" validate:"required"`
	PrivateKey   string      `json:"private_key" validate:"required"`
	Distribution string      `json:"distribution"`
	ConsumerAge  string      `json:"consumer_age"`
	ConsumerFreq string      `json:"consumer_freq"`
	Partnered    bool        `json:"partnered"`
	Timestamp    int64       `json:"timestamp"`
	Components   []Container `json:"components" validate:"required"`
}

type ComponentComponent struct {
	Items       *map[string]int `json:"items,omitempty"`
	Expiration  *bool           `json:"expiration,omitempty"`
	Benefits    *map[string]int `json:"benefits,omitempty"`
	Tier        *int64          `json:"tier,omitempty"`
	Paid        *bool           `json:"paid,omitempty"`
	Tasks       *map[string]int `json:"tasks,omitempty"`
	Distributor *string         `json:"distributor,omitempty"`
	Address     *string         `json:"address,omitempty"`
	District    *string         `json:"district,omitempty"`
	City        *string         `json:"city,omitempty"`
}
