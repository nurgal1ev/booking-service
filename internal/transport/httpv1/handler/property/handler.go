package property

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/nurgal1ev/booking-service/internal/service/property"
	"github.com/nurgal1ev/booking-service/internal/transport/middleware"
)

type PropertyHandler struct {
	propertyService *property.PropertyService
}

func NewPropertyHandler(propertyService *property.PropertyService) *PropertyHandler {
	return &PropertyHandler{
		propertyService: propertyService,
	}
}

func (h *PropertyHandler) CreateProperty(ctx context.Context, input *CreatePropertyInput) (*CreatePropertyOutput, error) {
	userID, ok := middleware.GetUserIDFromContext(ctx)
	if !ok || userID == 0 {
		return nil, huma.Error401Unauthorized("user not authenticated")
	}

	if input.Name == "" {
		return nil, huma.Error400BadRequest("name is required")
	}

	if input.PricePerNight <= 0 {
		return nil, huma.Error400BadRequest("price per night must be greater than 0")
	}

	if input.PropertyType == "" {
		return nil, huma.Error400BadRequest("property type is required")
	}

	property := &property.Property{
		Name:          input.Name,
		Description:   input.Description,
		Address:       input.Address,
		City:          input.City,
		Country:       input.Country,
		PricePerNight: input.PricePerNight,
		PropertyType:  input.PropertyType,
		OwnerID:       userID,
	}

	result, err := h.propertyService.Create(ctx, property)
	if err != nil {
		return nil, err
	}

	return &CreatePropertyOutput{
		ID:            result.ID,
		Name:          result.Name,
		Description:   result.Description,
		Address:       result.Address,
		City:          result.City,
		Country:       result.Country,
		PricePerNight: result.PricePerNight,
		PropertyType:  result.PropertyType,
	}, nil
}
