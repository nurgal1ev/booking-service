package unit

import (
	"context"
	"errors"
	"strings"

	"github.com/nurgal1ev/booking-service/internal/models"
	"github.com/nurgal1ev/booking-service/internal/repository/property"
	"github.com/nurgal1ev/booking-service/internal/repository/unit"
)

type UnitService struct {
	unitRepo     *unit.UnitRepo
	propertyRepo *property.PropertyRepo
}

func NewUnitService(unitRepo *unit.UnitRepo, propertyRepo *property.PropertyRepo) *UnitService {
	return &UnitService{
		unitRepo:     unitRepo,
		propertyRepo: propertyRepo,
	}
}

type Unit struct {
	ID            uint
	Name          string
	Description   string
	PricePerNight int
	Capacity      int
	IsAvailable   bool
}

type UpdateUnitRequest struct {
	Name          *string
	Description   *string
	PricePerNight *int
	Capacity      *int
	IsAvailable   *bool
}

func (u *Unit) Validate() error {
	if u.PricePerNight <= 0 {
		return errors.New("price per night must be greater than 0")
	}
	if u.Capacity <= 0 {
		return errors.New("capacity must be greater than 0")
	}
	if strings.TrimSpace(u.Name) == "" {
		return errors.New("name is required")
	}
	return nil
}

func (s *UnitService) Create(ctx context.Context, propertyID uint, userID uint, u *Unit) (*models.Unit, error) {
	property, err := s.propertyRepo.FindByID(ctx, propertyID)
	if err != nil {
		return nil, err
	}

	if property.OwnerID != userID {
		return nil, errors.New("you don't have permission to modify this unit")
	}

	if err := u.Validate(); err != nil {
		return nil, err
	}

	var unit = models.Unit{
		Name:          u.Name,
		Description:   u.Description,
		PricePerNight: u.PricePerNight,
		Capacity:      u.Capacity,
		IsAvailable:   u.IsAvailable,
		PropertyID:    propertyID,
	}

	err = s.unitRepo.Create(ctx, &unit)
	if err != nil {
		return nil, err
	}

	return &unit, nil
}

func (s *UnitService) Update(ctx context.Context, id uint, userID uint, req *UpdateUnitRequest) (*models.Unit, error) {
	unitId, err := s.unitRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if unitId == nil {
		return nil, errors.New("unit not found")
	}

	propertyID, err := s.propertyRepo.FindByID(ctx, unitId.PropertyID)
	if err != nil {
		return nil, err
	}

	if propertyID == nil {
		return nil, errors.New("property not found")
	}

	request := &unit.UpdateUnitRequest{
		Name:          req.Name,
		Description:   req.Description,
		PricePerNight: req.PricePerNight,
		Capacity:      req.Capacity,
		IsAvailable:   req.IsAvailable,
	}

	err = s.unitRepo.Update(ctx, id, request)
	if err != nil {
		return nil, err
	}

	return s.unitRepo.FindByID(ctx, id)
}
