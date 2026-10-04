package booking

import (
	"context"
	"errors"
	"time"

	"github.com/nurgal1ev/booking-service/internal/models"
	"github.com/nurgal1ev/booking-service/internal/repository/booking"
	"github.com/nurgal1ev/booking-service/internal/repository/unit"
)

type Booking struct {
	UnitID     uint
	CheckIn    time.Time
	CheckOut   time.Time
	GuestCount int
}

type BookingService struct {
	bookingRepo *booking.BookingRepo
	unitRepo    *unit.UnitRepo
}

func NewBookingService(bookingRepo *booking.BookingRepo, unitRepo *unit.UnitRepo) *BookingService {
	return &BookingService{bookingRepo: bookingRepo, unitRepo: unitRepo}
}

func (s *BookingService) Create(ctx context.Context, userID uint, b *Booking) (*models.Booking, error) {
	unit, err := s.unitRepo.FindByID(ctx, b.UnitID)
	if err != nil {
		return nil, err
	}

	if unit == nil {
		return nil, errors.New("unit not found")
	}

	if !b.CheckIn.Before(b.CheckOut) {
		return nil, errors.New("check_in must be before check_out")
	}

	duration := b.CheckOut.Sub(b.CheckIn)
	nights := int(duration.Hours() / 24)
	totalPrice := nights * unit.PricePerNight

	if b.GuestCount > unit.Capacity {
		return nil, errors.New("too many guests for this unit")
	}

	var booking = models.Booking{
		UserID:     userID,
		UnitID:     b.UnitID,
		CheckIn:    b.CheckIn,
		CheckOut:   b.CheckOut,
		GuestCount: b.GuestCount,
		TotalPrice: totalPrice,
		Status:     "pending",
	}

	err = s.bookingRepo.CreateWithLock(ctx, &booking)
	if err != nil {
		return nil, err
	}

	return &booking, nil
}

func (s *BookingService) GetUserBookings(ctx context.Context, userID uint) ([]models.Booking, error) {
	bookings, err := s.bookingRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return bookings, nil
}
