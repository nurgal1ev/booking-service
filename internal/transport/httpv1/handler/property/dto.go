package property

type Property struct {
	ID            uint               `json:"id"`
	Name          string             `json:"name,omitempty" minLength:"1" maxLength:"255" pattern:"^[a-zA-Zа-яА-Я0-9\\s]+$"`
	Description   string             `json:"description,omitempty" maxLength:"10000" pattern:"^[\\p{L}\\p{N}\\p{P}\\p{Z}\\n\\r]+$"`
	Address       string             `json:"address"`
	City          string             `json:"city"`
	Country       string             `json:"country"`
	PricePerNight int                `json:"price_per_night"`
	PropertyType  string             `json:"property_type"`
	Units         []*PropertyUnitDTO `json:"units"`
}

type PropertyUnitDTO struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	PricePerNight int    `json:"price_per_night"`
	Capacity      int    `json:"capacity"`
	IsAvailable   bool   `json:"is_available"`
}

type CreatePropertyInput struct {
	Body Property
}

type CreatePropertyOutput struct {
	Body Property
}

type GetPropertyInput struct {
	ID uint `path:"id"`
}

type GetPropertyOutput struct {
	Body Property
}

type UpdatePropertyInput struct {
	ID   uint `path:"id"`
	Body struct {
		Name          *string `json:"name,omitempty" minLength:"1" maxLength:"255" pattern:"^[a-zA-Zа-яА-Я0-9\\s]+$"`
		Description   *string `json:"description,omitempty" maxLength:"10000" pattern:"^[\\p{L}\\p{N}\\p{P}\\p{Z}\\n\\r]+$"`
		Address       *string `json:"address"`
		City          *string `json:"city"`
		Country       *string `json:"country"`
		PricePerNight *int    `json:"price_per_night"`
		PropertyType  *string `json:"property_type"`
	}
}

type UpdatePropertyOutput struct {
	Body Property
}

type DeletePropertyInput struct {
	ID uint `path:"id"`
}

type DeletePropertyOutput struct {
	Body struct {
		Message string `json:"message"`
	}
}
