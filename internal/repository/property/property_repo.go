package property

import (
	"context"
	"errors"

	"github.com/nurgal1ev/booking-service/internal/models"
	"gorm.io/gorm"
)

type PropertyRepo struct {
	db *gorm.DB
}

func NewPropertyRepo(db *gorm.DB) *PropertyRepo {
	return &PropertyRepo{db: db}
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

func (p *PropertyRepo) Create(ctx context.Context, property *models.Property) error {
	query := p.db.WithContext(ctx).Create(property).Error
	return query
}

func (p *PropertyRepo) FindByID(ctx context.Context, id uint) (*models.Property, error) {
	var property models.Property
	err := p.db.WithContext(ctx).Where("id = ?", id).First(&property).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &property, nil
}

func (p *PropertyRepo) FindAll(ctx context.Context, filter PropertyFilter) ([]models.Property, error) {
	query := p.db.WithContext(ctx).Model(&models.Property{})

	if filter.City != "" {
		query = query.Where("city = ?", filter.City)
	}
	if filter.PropertyType != "" {
		query = query.Where("property_type = ?", filter.PropertyType)
	}
	if filter.MinPrice > 0 {
		query = query.Where("price_per_night >= ?", filter.MinPrice)
	}
	if filter.MaxPrice > 0 {
		query = query.Where("price_per_night <= ?", filter.MaxPrice)
	}
	if filter.Search != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Search+"%")
	}

	var properties []models.Property
	err := query.Find(&properties).Error
	return properties, err
}

func (p *PropertyRepo) FindByOwnerID(ctx context.Context, id uint) ([]models.Property, error) {
	var properties []models.Property
	err := p.db.WithContext(ctx).Where("owner_id = ?", id).Find(&properties).Error
	if err != nil {
		return nil, err
	}

	return properties, nil
}

func (p *PropertyRepo) Update(ctx context.Context, id uint, req *UpdatePropertyRequest) error {
	updates := make(map[string]interface{})

	if req.Name != nil {
		updates["name"] = *req.Name
	}

	if req.Description != nil {
		updates["description"] = *req.Description
	}

	if req.Address != nil {
		updates["address"] = *req.Address
	}

	if req.City != nil {
		updates["city"] = *req.City
	}

	if req.Country != nil {
		updates["country"] = *req.Country
	}

	if req.PricePerNight != nil {
		updates["price_per_night"] = *req.PricePerNight
	}

	return p.db.WithContext(ctx).Model(&models.Property{}).Where("id = ?", id).Updates(updates).Error
}

func (p *PropertyRepo) Delete(ctx context.Context, id uint) error {
	result := p.db.WithContext(ctx).Delete(&models.Property{}, id)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("property not found")
	}

	return nil
}
