package domain

type Quote struct {
	Carrier []Carrier `json:"carrier"`
}

type Carrier struct {
	Name     string  `json:"name"`
	Service  string  `json:"service"`
	Deadline uint    `json:"deadline"`
	Price    float64 `json:"price"`
}
