package postgres

import (
	"log"

	"github.com/nurgal1ev/booking-service/internal/config"
	"github.com/nurgal1ev/booking-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(c *config.Config) *Db {
	db, err := gorm.Open(postgres.Open(c.DbConfig.DSN()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		models.User{},
		models.Properties{},
		models.Booking{},
	)

	if err != nil {
		log.Printf("Error during migration: %v", err)
		return nil
	}

	return &Db{db}
}
