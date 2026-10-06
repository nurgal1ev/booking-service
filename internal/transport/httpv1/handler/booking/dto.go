package booking

import "time"

type Booking struct {
	ID         uint      `json:"id"`
	UserID     uint      `json:"user_id"`
	UnitID     uint      `json:"unit_id"`
	CheckIn    time.Time `json:"check_in"`
	CheckOut   time.Time `json:"check_out"`
	TotalPrice int       `json:"total_price"`
	Status     string    `json:"status"`
	GuestCount int       `json:"guest_count"`
}

type CreateBookingInput struct {
	Body struct {
		UnitID     uint      `json:"unit_id"`
		CheckIn    time.Time `json:"check_in"`
		CheckOut   time.Time `json:"check_out"`
		GuestCount int       `json:"guest_count"`
	}
}

type CreateBookingOutput struct {
	Body Booking
}
