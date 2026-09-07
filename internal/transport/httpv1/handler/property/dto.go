package property

type CreatePropertyInput struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	Address       string `json:"address"`
	City          string `json:"city"`
	Country       string `json:"country"`
	PricePerNight int    `json:"price_per_night"`
	PropertyType  string `json:"property_type"`
}

type CreatePropertyOutput struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Address       string `json:"address"`
	City          string `json:"city"`
	Country       string `json:"country"`
	PricePerNight int    `json:"price_per_night"`
	PropertyType  string `json:"property_type"`
}
