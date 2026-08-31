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

type PropertyFilter struct {
	City         string
	Country      string
	PropertyType string
	MinPrice     int
	MaxPrice     int
	Search       string
	OwnerID      uint
}

type UpdatePropertyRequest struct {
	Name          *string
	Description   *string
	Address       *string
	City          *string
	Country       *string
	PricePerNight *int
	PropertyType  *string
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

func (s *PropertyService) GetById(ctx context.Context, id uint) (*models.Property, error) {
	if id == 0 {
		return nil, errors.New("invalid id")
	}

	property, err := s.propertyRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if property == nil {
		return nil, err
	}

	return property, nil
}

func (s *PropertyService) GetAll(ctx context.Context, filter PropertyFilter) ([]models.Property, error) {
	if filter.MinPrice > filter.MaxPrice {
		return nil, errors.New("minPrice cannot be greater than maxPrice")
	}

	if filter.PropertyType != "" {
		validTypes := []string{"hotel", "apartment", "house", "camping"}
		isValid := false
		for _, t := range validTypes {
			if filter.PropertyType == t {
				isValid = true
				break
			}
		}
		if !isValid {
			return nil, errors.New("invalid property type. Allowed: hotel, apartment, house, camping")
		}
	}

	result := property.PropertyFilter{
		City:         filter.City,
		Country:      filter.Country,
		PropertyType: filter.PropertyType,
		MinPrice:     filter.MinPrice,
		MaxPrice:     filter.MaxPrice,
		Search:       filter.Search,
	}

	properties, err := s.propertyRepo.FindAll(ctx, result)
	if err != nil {
		return nil, err
	}

	return properties, nil
}

func (s *PropertyService) Update(ctx context.Context, id uint, userID uint, userRole string, req *UpdatePropertyRequest) (*models.Property, error) {
	propertyId, err := s.propertyRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if propertyId == nil {
		return nil, errors.New("property not found")
	}

	if propertyId.OwnerID != userID && userRole != "admin" {
		return nil, errors.New("you don't have permission to update this property")
	}

	request := &property.UpdatePropertyRequest{
		Name:          req.Name,
		Description:   req.Description,
		Address:       req.Address,
		City:          req.City,
		Country:       req.Country,
		PricePerNight: req.PricePerNight,
		PropertyType:  req.PropertyType,
	}

	err = s.propertyRepo.Update(ctx, id, request)
	if err != nil {
		return nil, err
	}

	return s.propertyRepo.FindByID(ctx, id)
}

func (s *PropertyService) Delete(ctx context.Context, id uint, userID uint, userRole string) error {
	propertyId, err := s.propertyRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if propertyId == nil {
		return errors.New("property not found")
	}

	if propertyId.OwnerID != userID && userRole != "admin" {
		return errors.New("you don't have permission to delete this property")
	}

	err = s.propertyRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
