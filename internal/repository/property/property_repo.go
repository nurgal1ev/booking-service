package property

import (
	"context"

	"github.com/nurgal1ev/booking-service/internal/models"
	"gorm.io/gorm"
)

type PropertyRepo struct {
	db *gorm.DB
}

func NewPropertyRepo(db *gorm.DB) *PropertyRepo {
	return &PropertyRepo{db: db}
}

func (p *PropertyRepo) Create(ctx context.Context, property *models.Property) error {
	query := p.db.WithContext(ctx).Create(property).Error
	return query
}
