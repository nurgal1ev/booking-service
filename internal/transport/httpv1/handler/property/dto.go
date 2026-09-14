package property

type CreatePropertyInput struct {
	Body struct {
		Name          string `json:"name"`
		Description   string `json:"description"`
		Address       string `json:"address"`
		City          string `json:"city"`
		Country       string `json:"country"`
		PricePerNight int    `json:"price_per_night"`
		PropertyType  string `json:"property_type"`
	}
}

type CreatePropertyOutput struct {
	Body struct {
		ID            uint   `json:"id"`
		Name          string `json:"name"`
		Description   string `json:"description"`
		Address       string `json:"address"`
		City          string `json:"city"`
		Country       string `json:"country"`
		PricePerNight int    `json:"price_per_night"`
		PropertyType  string `json:"property_type"`
	}
}

type GetPropertyInput struct {
	ID uint `path:"id"`
}

type GetPropertyOutput struct {
	Body struct {
		ID            uint   `json:"id"`
		Name          string `json:"name"`
		Description   string `json:"description"`
		Address       string `json:"address"`
		City          string `json:"city"`
		Country       string `json:"country"`
		PricePerNight int    `json:"price_per_night"`
		PropertyType  string `json:"property_type"`
	}
}

type UpdatePropertyInput struct {
	ID   uint `path:"id"`
	Body struct {
		Name          *string `json:"Name"`
		Description   *string `json:"description"`
		Address       *string `json:"address"`
		City          *string `json:"city"`
		Country       *string `json:"country"`
		PricePerNight *int    `json:"price_per_night"`
		PropertyType  *string `json:"property_type"`
	}
}

type UpdatePropertyOutput struct {
	Body struct {
		Name          string
		Description   string
		Address       string
		City          string
		Country       string
		PricePerNight int
		PropertyType  string
	}
}
