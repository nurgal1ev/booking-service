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
