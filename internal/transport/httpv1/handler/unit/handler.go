package unit

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/nurgal1ev/booking-service/internal/service/unit"
	"github.com/nurgal1ev/booking-service/internal/transport/middleware"
)

type UnitHandler struct {
	unitService *unit.UnitService
}

func NewUnitHandler(unitService *unit.UnitService) *UnitHandler {
	return &UnitHandler{
		unitService: unitService,
	}
}

func (h *UnitHandler) CreateUnitHandler(ctx context.Context, input *CreateUnitInput) (*CreateUnitOutput, error) {
	userID, ok := middleware.GetUserIDFromContext(ctx)
	if !ok || userID == 0 {
		return nil, huma.Error401Unauthorized("user not authenticated")
	}

	unit := unit.Unit{
		Name:          input.Body.Name,
		Description:   input.Body.Description,
		Capacity:      input.Body.Capacity,
		PricePerNight: input.Body.PricePerNight,
	}

	result, err := h.unitService.Create(ctx, input.PropertyID, userID, &unit)
	if err != nil {
		return nil, err
	}

	return &CreateUnitOutput{
		Body: struct {
			ID            uint   `json:"id"`
			Name          string `json:"name"`
			Description   string `json:"description"`
			Capacity      int    `json:"capacity"`
			PricePerNight int    `json:"price_per_night"`
		}{
			ID:            result.ID,
			Name:          result.Name,
			Description:   result.Description,
			Capacity:      result.Capacity,
			PricePerNight: result.PricePerNight,
		},
	}, nil
}
