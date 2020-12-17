package model

type Metadata struct {
	Type  string `json:"type"`
	Label string `json:"label"`
}

func NewMetadata() *Metadata {
	return &Metadata{Type: "external"}
}
