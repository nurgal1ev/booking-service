package postgres

import (
	"github.com/nurgal1ev/booking-service/internal/config"
	"github.com/nurgal1ev/booking-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(c *config.Config) (*Db, error) {
	db, err := gorm.Open(postgres.Open(c.DbConfig.DSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		models.User{},
		models.Property{},
		models.Booking{},
	)
	if err != nil {
		return nil, err
	}

	return &Db{
		DB: db,
	}, nil
}
