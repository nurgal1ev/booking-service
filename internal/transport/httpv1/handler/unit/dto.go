package unit

type UnitDTO struct {
	ID            uint
	Name          string
	Description   string
	PricePerNight int
	Capacity      int
	IsAvailable   bool
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
	PropertyID uint `json:"property_id"`
	Body       []UnitDTO
}
