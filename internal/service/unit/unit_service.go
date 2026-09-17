package unit

import (
	"context"

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

func (s *UnitService) Create(ctx context.Context, u *Unit) (*models.Unit, error) {
	return nil, err
}
