package unit

type CreateUnitInput struct {
	PropertyID uint `path:"id"`
	Body       struct {
		Name          string `json:"name"`
		Description   string `json:"description"`
		Capacity      int    `json:"capacity"`
		PricePerNight int    `json:"price_per_night"`
	}
}

type CreateUnitOutput struct {
	Body struct {
		ID            uint   `json:"id"`
		Name          string `json:"name"`
		Description   string `json:"description"`
		Capacity      int    `json:"capacity"`
		PricePerNight int    `json:"price_per_night"`
	}
}
