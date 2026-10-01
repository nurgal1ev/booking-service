package unit

type Unit struct {
	ID            uint   `json:"id"`
	Name          string `json:"name,omitempty" minLength:"1" maxLength:"55" pattern:"^[a-zA-Zа-яА-Я0-9\\s]+$"`
	Description   string `json:"description,omitempty" maxLength:"10000" pattern:"^[\\p{L}\\p{N}\\p{P}\\p{Z}\\n\\r]+$"`
	PricePerNight int    `json:"price_per_night"`
	Capacity      int    `json:"capacity"`
	IsAvailable   bool   `json:"is_available"`
}

type CreateUnitInput struct {
	PropertyID uint `path:"id"`
	Body       struct {
		Name          string `json:"name,omitempty" minLength:"1" maxLength:"55" pattern:"^[a-zA-Zа-яА-Я0-9\\s]+$"`
		Description   string `json:"description"`
		Capacity      int    `json:"capacity"`
		PricePerNight int    `json:"price_per_night"`
	}
}

type CreateUnitOutput struct {
	Body Unit
}

type GetUnitInput struct {
	PropertyID uint `path:"id"`
}

type GetUnitOutput struct {
	Body []*Unit
}

type UpdateUnitInput struct {
	ID   uint `path:"id"`
	Body struct {
		Name          *string `json:"name,omitempty" minLength:"1" maxLength:"55" pattern:"^[a-zA-Zа-яА-Я0-9\\s]+$"`
		Description   *string `json:"description,omitempty" maxLength:"10000" pattern:"^[\\p{L}\\p{N}\\p{P}\\p{Z}\\n\\r]+$"`
		Capacity      *int    `json:"capacity"`
		PricePerNight *int    `json:"price_per_night"`
		IsAvailable   *bool   `json:"is_available"`
	}
}

type UpdateUnitOutput struct {
	Body Unit
}

type DeleteUnitInput struct {
	ID uint `path:"id"`
}

type DeleteUnitOutput struct {
	Body struct {
		Message string `json:"message"`
	}
}
