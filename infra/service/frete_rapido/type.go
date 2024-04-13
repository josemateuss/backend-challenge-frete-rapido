package frete_rapido

type Service struct{}

type RequestPayload struct {
	Shipper        Shipper      `json:"shipper"`
	Recipient      Recipient    `json:"recipient"`
	Dispatchers    []Dispatcher `json:"dispatchers"`
	SimulationType []uint       `json:"simulation_type"`
}

type Shipper struct {
	RegisteredNumber string `json:"registered_number"`
	Token            string `json:"token"`
	PlatformCode     string `json:"platform_code"`
}

type Recipient struct {
	Zipcode uint `json:"zipcode"`
}

type Dispatcher struct {
	RegisteredNumber string   `json:"registered_number"`
	Zipcode          uint     `json:"zipcode"`
	Volumes          []Volume `json:"volumes"`
}

type Volume struct {
	Amount        uint    `json:"amount"`
	AmountVolumes uint    `json:"amount_volumes"`
	Category      string  `json:"category"`
	Sku           string  `json:"sku"`
	Height        float64 `json:"height"`
	Width         float64 `json:"width"`
	Length        float64 `json:"length"`
	UnitaryPrice  float64 `json:"unitary_price"`
	UnitaryWeight float64 `json:"unitary_weight"`
}

type ResponsePayload struct {
	Dispatchers []ResponseDispatcher `json:"dispatchers"`
}

type ResponseDispatcher struct {
	Offers []Offer `json:"offers"`
}

type Offer struct {
	Carrier      Carrier      `json:"carrier"`
	Service      string       `json:"service"`
	DeliveryTime DeliveryTime `json:"delivery_time"`
	FinalPrice   float64      `json:"final_price"`
}

type Carrier struct {
	Name string `json:"name"`
}

type DeliveryTime struct {
	Days uint `json:"days"`
}
