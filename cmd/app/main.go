package main

import (
	"log"

	"github.com/nurgal1ev/booking-service/internal/config"
	"github.com/nurgal1ev/booking-service/internal/infrastructure/postgres"
)

func main() {
	cfg := config.Load()
	db := postgres.NewDb(cfg)
	if db == nil {
		log.Fatal("Failed to initialize database")
	}
}
