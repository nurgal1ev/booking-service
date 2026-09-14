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

func (h *PropertyHandler) CreatePropertyHandler(ctx context.Context, input *CreatePropertyInput) (*CreatePropertyOutput, error) {
	userID, ok := middleware.GetUserIDFromContext(ctx)
	if !ok || userID == 0 {
		return nil, huma.Error401Unauthorized("user not authenticated")
	}

	if input.Body.Name == "" {
		return nil, huma.Error400BadRequest("name is required")
	}

	if input.Body.PricePerNight <= 0 {
		return nil, huma.Error400BadRequest("price per night must be greater than 0")
	}

	if input.Body.PropertyType == "" {
		return nil, huma.Error400BadRequest("property type is required")
	}

	property := &property.Property{
		Name:          input.Body.Name,
		Description:   input.Body.Description,
		Address:       input.Body.Address,
		City:          input.Body.City,
		Country:       input.Body.Country,
		PricePerNight: input.Body.PricePerNight,
		PropertyType:  input.Body.PropertyType,
		OwnerID:       userID,
	}

	result, err := h.propertyService.Create(ctx, property)
	if err != nil {
		return nil, err
	}

	return &CreatePropertyOutput{
		Body: struct {
			ID            uint   `json:"id"`
			Name          string `json:"name"`
			Description   string `json:"description"`
			Address       string `json:"address"`
			City          string `json:"city"`
			Country       string `json:"country"`
			PricePerNight int    `json:"price_per_night"`
			PropertyType  string `json:"property_type"`
		}{
			ID:            result.ID,
			Name:          result.Name,
			Description:   result.Description,
			Address:       result.Address,
			City:          result.City,
			Country:       result.Country,
			PricePerNight: result.PricePerNight,
			PropertyType:  result.PropertyType,
		},
	}, nil
}

func (h *PropertyHandler) GetPropertyHandler(ctx context.Context, input *GetPropertyInput) (*GetPropertyOutput, error) {
	property, err := h.propertyService.GetById(ctx, input.ID)
	if err != nil {
		return nil, huma.Error404NotFound("property not found")
	}

	if property == nil {
		return nil, huma.Error404NotFound("property not found")
	}

	return &GetPropertyOutput{
		Body: struct {
			ID            uint   `json:"id"`
			Name          string `json:"name"`
			Description   string `json:"description"`
			Address       string `json:"address"`
			City          string `json:"city"`
			Country       string `json:"country"`
			PricePerNight int    `json:"price_per_night"`
			PropertyType  string `json:"property_type"`
		}{
			ID:            property.ID,
			Name:          property.Name,
			Description:   property.Description,
			Address:       property.Address,
			City:          property.City,
			Country:       property.Country,
			PricePerNight: property.PricePerNight,
			PropertyType:  property.PropertyType,
		},
	}, nil
}

func (h *PropertyHandler) UpdatePropertyHandler(ctx context.Context, input *UpdatePropertyInput) (*UpdatePropertyOutput, error) {
	propertyID, err := h.propertyService.GetById(ctx, input.ID)
	if err != nil {
		return nil, huma.Error404NotFound("property not found")
	}

	userID, ok := middleware.GetUserIDFromContext(ctx)
	if !ok || userID == 0 {
		return nil, huma.Error401Unauthorized("user not authenticated")
	}

	userRole, _ := middleware.GetUserRoleFromContext(ctx)

	if propertyID.OwnerID != userID {
		return nil, huma.Error403Forbidden("forbidden")
	}

	req := &property.UpdatePropertyRequest{
		Name:          input.Body.Name,
		Description:   input.Body.Description,
		Address:       input.Body.Address,
		City:          input.Body.City,
		Country:       input.Body.Country,
		PricePerNight: input.Body.PricePerNight,
		PropertyType:  input.Body.PropertyType,
	}

	updatedProperty, err := h.propertyService.Update(ctx, input.ID, userID, userRole, req)
	if err != nil {
		return nil, huma.Error400BadRequest(err.Error())
	}

	return &UpdatePropertyOutput{
		Body: struct {
			Name          string
			Description   string
			Address       string
			City          string
			Country       string
			PricePerNight int
			PropertyType  string
		}{
			Name:          updatedProperty.Name,
			Description:   updatedProperty.Description,
			Address:       updatedProperty.Address,
			City:          updatedProperty.City,
			Country:       updatedProperty.Country,
			PricePerNight: updatedProperty.PricePerNight,
			PropertyType:  updatedProperty.PropertyType,
		},
	}, nil
}

func (h *PropertyHandler) DeletePropertyHandler(ctx context.Context, input *DeletePropertyInput) (*DeletePropertyOutput, error) {
	propertyID, err := h.propertyService.GetById(ctx, input.ID)
	if err != nil {
		return nil, huma.Error404NotFound("property not found")
	}

	userID, ok := middleware.GetUserIDFromContext(ctx)
	if !ok || userID == 0 {
		return nil, huma.Error401Unauthorized("user not authenticated")
	}

	userRole, _ := middleware.GetUserRoleFromContext(ctx)

	err = h.propertyService.Delete(ctx, propertyID.ID, userID, userRole)
	if err != nil {
		return nil, err
	}

	return &DeletePropertyOutput{
		Body: struct {
			Message string `json:"message"`
		}{
			Message: "property deleted successfully",
		},
	}, nil
}
