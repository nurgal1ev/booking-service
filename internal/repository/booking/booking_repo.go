package booking

import (
	"context"

	"github.com/nurgal1ev/booking-service/internal/models"
	"gorm.io/gorm"
)

type BookingRepo struct {
	db *gorm.DB
}

func NewBookingRepo(db *gorm.DB) *BookingRepo {
	return &BookingRepo{db: db}
}

func (r *BookingRepo) Create(ctx context.Context, booking models.Booking) error {
	query := r.db.WithContext(ctx).Create(&booking).Error
	return query
}
