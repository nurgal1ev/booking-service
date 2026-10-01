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
		Body: Unit{
			Name:          result.Name,
			Description:   result.Description,
			Capacity:      result.Capacity,
			PricePerNight: result.PricePerNight,
		},
	}, nil
}

func (h *UnitHandler) GetUnitByPropertyIdHandler(ctx context.Context, input *GetUnitInput) (*GetUnitOutput, error) {
	units, err := h.unitService.GetByPropertyID(ctx, input.PropertyID)
	if err != nil {
		return nil, huma.Error404NotFound("unit not found")
	}

	unitDTOs := make([]*Unit, len(units))
	for i, unit := range units {
		unitDTOs[i] = &Unit{
			ID:            unit.ID,
			Name:          unit.Name,
			Description:   unit.Description,
			Capacity:      unit.Capacity,
			PricePerNight: unit.PricePerNight,
			IsAvailable:   unit.IsAvailable,
		}
	}

	return &GetUnitOutput{
		Body: unitDTOs,
	}, nil
}

func (h *UnitHandler) UpdateUnitHandler(ctx context.Context, input *UpdateUnitInput) (*UpdateUnitOutput, error) {
	userID, ok := middleware.GetUserIDFromContext(ctx)
	if !ok || userID == 0 {
		return nil, huma.Error401Unauthorized("user not authenticated")
	}

	req := &unit.UpdateUnitRequest{
		Name:          input.Body.Name,
		Description:   input.Body.Description,
		Capacity:      input.Body.Capacity,
		PricePerNight: input.Body.PricePerNight,
		IsAvailable:   input.Body.IsAvailable,
	}

	updatedUnit, err := h.unitService.Update(ctx, input.ID, userID, req)
	if err != nil {
		return nil, huma.Error400BadRequest(err.Error())
	}

	return &UpdateUnitOutput{
		Body: Unit{
			ID:            updatedUnit.ID,
			Name:          updatedUnit.Name,
			Description:   updatedUnit.Description,
			PricePerNight: updatedUnit.PricePerNight,
			Capacity:      updatedUnit.Capacity,
			IsAvailable:   updatedUnit.IsAvailable,
		},
	}, nil
}

func (h *UnitHandler) DeleteUnitHandler(ctx context.Context, input *DeleteUnitInput) (*DeleteUnitOutput, error) {
	userID, ok := middleware.GetUserIDFromContext(ctx)
	if !ok || userID == 0 {
		return nil, huma.Error401Unauthorized("user not authenticated")
	}

	userRole, _ := middleware.GetUserRoleFromContext(ctx)

	err := h.unitService.Delete(ctx, input.ID, userID, userRole)
	if err != nil {
		return nil, err
	}

	return &DeleteUnitOutput{
		Body: struct {
			Message string `json:"message"`
		}{
			Message: "unit deleted successfully",
		},
	}, nil
}
