package unit

import (
	"context"
	"errors"

	"github.com/nurgal1ev/booking-service/internal/models"
	"gorm.io/gorm"
)

type UnitRepo struct {
	db *gorm.DB
}

func NewUnitRepo(db *gorm.DB) *UnitRepo {
	return &UnitRepo{db: db}
}

type UpdateUnitRequest struct {
	Name          *string
	Description   *string
	PricePerNight *int
	Capacity      *int
	IsAvailable   *bool
}

func (u *UnitRepo) Create(ctx context.Context, unit *models.Unit) error {
	query := u.db.WithContext(ctx).Create(unit).Error
	return query
}

func (u *UnitRepo) FindByID(ctx context.Context, id uint) (*models.Unit, error) {
	var unit models.Unit
	err := u.db.WithContext(ctx).Where("id = ?", id).First(&unit).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &unit, nil
}

func (u *UnitRepo) FindByPropertyID(ctx context.Context, propertyID uint) ([]*models.Unit, error) {
	var units []*models.Unit
	err := u.db.WithContext(ctx).Where("property_id = ?", propertyID).Find(&units).Error
	if err != nil {
		return nil, err
	}
	return units, err
}

func (u *UnitRepo) Update(ctx context.Context, id uint, req *UpdateUnitRequest) error {
	updates := make(map[string]interface{})

	if req.Name != nil {
		updates["name"] = *req.Name
	}

	if req.Description != nil {
		updates["description"] = *req.Description
	}

	if req.PricePerNight != nil && *req.PricePerNight > 0 {
		updates["price_per_night"] = *req.PricePerNight
	}

	if req.Capacity != nil && *req.Capacity > 0 {
		updates["capacity"] = *req.Capacity
	}

	if req.IsAvailable != nil {
		updates["is_available"] = *req.IsAvailable
	}

	return u.db.WithContext(ctx).Model(&models.Unit{}).Where("id = ?", id).Updates(updates).Error
}

func (u *UnitRepo) Delete(ctx context.Context, id uint) error {
	result := u.db.WithContext(ctx).Delete(&models.Unit{}, id)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("unit not found")
	}

	return nil
}
