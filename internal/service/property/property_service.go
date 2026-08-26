package property

import (
	"context"
	"errors"
	"strings"

	"github.com/nurgal1ev/booking-service/internal/models"
	"github.com/nurgal1ev/booking-service/internal/repository/property"
	"github.com/nurgal1ev/booking-service/internal/repository/user"
)

type Property struct {
	ID            uint
	Name          string
	Description   string
	Address       string
	City          string
	Country       string
	PricePerNight int
	PropertyType  string
	OwnerID       uint
}

type PropertyService struct {
	propertyRepo *property.PropertyRepo
	userRepo     *user.UserRepo
}

func NewPropertyService(propertyRepo *property.PropertyRepo, userRepo *user.UserRepo) *PropertyService {
	return &PropertyService{
		propertyRepo: propertyRepo,
		userRepo:     userRepo,
	}
}

func (p *Property) Validate() error {
	if strings.TrimSpace(p.Name) == "" {
		return errors.New("property name cannot be empty")
	}

	if p.PricePerNight <= 0 {
		return errors.New("price per night must be greater than zero")
	}

	if p.OwnerID == 0 {
		return errors.New("owner id is required")
	}

	validTypes := []string{"hotel", "apartment", "house", "villa"}
	for _, t := range validTypes {
		if p.PropertyType == t {
			return nil
		}
	}
	return errors.New("property type must be: hotel, apartment, house, or villa")
}

func (s *PropertyService) Create(ctx context.Context, p *Property) (*models.Property, error) {
	user, err := s.userRepo.FindByID(ctx, p.OwnerID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	var property = models.Property{
		Name:          p.Name,
		Description:   p.Description,
		Address:       p.Address,
		City:          p.City,
		Country:       p.Country,
		PricePerNight: p.PricePerNight,
		PropertyType:  p.PropertyType,
		OwnerID:       p.OwnerID,
	}

	err = s.propertyRepo.Create(ctx, &property)
	if err != nil {
		return nil, err
	}

	return &property, nil
}
