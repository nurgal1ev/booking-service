package postgres

import (
	"github.com/nurgal1ev/booking-service/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(c *config.Config) *Db {
	db, err := gorm.Open(postgres.Open(c.DbConfig.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &Db{db}
}
