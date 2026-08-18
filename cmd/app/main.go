package main

import (
	"log"

	"github.com/nurgal1ev/booking-service/internal/app"
	"github.com/nurgal1ev/booking-service/internal/config"
	"github.com/nurgal1ev/booking-service/internal/infrastructure/postgres"
	"github.com/nurgal1ev/booking-service/internal/transport/httpv1"
)

func main() {
	cfg := config.Load()

	db, err := postgres.NewDb(cfg)
	if err != nil {
		log.Fatal(err)
	}

	handlers := app.InitHandlers(db.DB, cfg)

	httpv1.StartServer(handlers)
}
