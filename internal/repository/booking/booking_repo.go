package booking

import (
	"context"
	"errors"
	"time"

	"github.com/nurgal1ev/booking-service/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (r *BookingRepo) FindByID(ctx context.Context, id uint) (*models.Booking, error) {
	var booking models.Booking
	err := r.db.WithContext(ctx).First(&booking, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &booking, nil
}

func (r *BookingRepo) FindOverlapping(ctx context.Context, tx *gorm.DB, unitID uint, checkIn, checkOut time.Time) (bool, error) {
	var count int64

	err := tx.WithContext(ctx).
		Model(&models.Booking{}).
		Where("unit_id = ?", unitID).
		Where("status != ?", "cancelled").
		Where("check_in < ?", checkOut).
		Where("check_out > ?", checkIn).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *BookingRepo) CreateWithLock(ctx context.Context, booking *models.Booking) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var unit models.Unit
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&unit, booking.UnitID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("unit not found")
			}
			return err
		}

		overlapping, err := r.FindOverlapping(ctx, tx, booking.UnitID, booking.CheckIn, booking.CheckOut)
		if err != nil {
			return err
		}
		if overlapping {
			return errors.New("unit is already booked for these dates")
		}

		if err := tx.Create(booking).Error; err != nil {
			return err
		}

		return nil
	})
}
