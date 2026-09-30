package unit

type Unit struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	PricePerNight int    `json:"price_per_night"`
	Capacity      int    `json:"capacity"`
	IsAvailable   bool   `json:"is_available"`
}

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

type GetUnitInput struct {
	PropertyID uint `path:"id"`
}

type GetUnitOutput struct {
	Body []*Unit
}

type UpdateUnitInput struct {
	ID   uint `path:"ID"`
	Body struct {
		Name          *string `json:"name"`
		Description   *string `json:"description"`
		Capacity      *int    `json:"capacity"`
		PricePerNight *int    `json:"price_per_night"`
		IsAvailable   *bool   `json:"is_availabale"`
	}
}

type UpdateUnitOutput struct {
	Body Unit
}
