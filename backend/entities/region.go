package entities

type Region struct {
	Bundesland string `json:"bundesland"`
	Kreis      string `json:"kreis"`
	Stadt      string `json:"stadt"`
}
