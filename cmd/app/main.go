package main

import (
	"github.com/nurgal1ev/booking-service/internal/config"
	"github.com/nurgal1ev/booking-service/internal/infrastructure/postgres"
)

func main() {
	cfg := config.Load()
	postgres.NewDb(cfg)
}
