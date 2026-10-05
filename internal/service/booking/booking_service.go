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

type UpdateBookingRequest struct {
	Status *string
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

func (s *BookingService) Cancel(ctx context.Context, id uint, userID uint) (*models.Booking, error) {
	bookingId, err := s.bookingRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if bookingId == nil {
		return nil, errors.New("booking not found")
	}

	if bookingId.UserID != userID {
		return nil, errors.New("user is not authorized to cancel this booking")
	}

	if bookingId.Status != "pending" && bookingId.Status != "confirmed" {
		return nil, errors.New("cannot cancel this booking")
	}

	err = s.bookingRepo.UpdateStatus(ctx, id, "cancelled")
	if err != nil {
		return nil, err
	}

	return s.bookingRepo.FindByID(ctx, id)
}

func (s *BookingService) Confirm(ctx context.Context, id uint, userRole string) error {
	bookingId, err := s.bookingRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if bookingId == nil {
		return errors.New("booking not found")
	}

	if userRole != "admin" {
		return errors.New("user is not authorized to confirm this booking")
	}

	if bookingId.Status != "pending" {
		return errors.New("cannot confirm this booking")
	}

	err = s.bookingRepo.UpdateStatus(ctx, id, "confirmed")
	if err != nil {
		return err
	}

	return nil
}
