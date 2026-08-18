package main

import (
	"log"

	"github.com/nurgal1ev/booking-service/internal/config"
	"github.com/nurgal1ev/booking-service/internal/infrastructure/postgres"
	userRepo "github.com/nurgal1ev/booking-service/internal/repository/user"
	userService "github.com/nurgal1ev/booking-service/internal/service/user"
	httpv1 "github.com/nurgal1ev/booking-service/internal/transport/httpv1"
	userHandler "github.com/nurgal1ev/booking-service/internal/transport/httpv1/handler/user"
)

func main() {
	cfg := config.Load()

	db, err := postgres.NewDb(cfg)
	if err != nil {
		log.Fatal(err)
	}

	userRepository := userRepo.NewUserRepo(db.DB)
	userService := userService.NewUserService(userRepository)
	userHandler := userHandler.NewUserHandler(userService)

	handlers := httpv1.Handlers{
		User: userHandler,
	}

	httpv1.StartServer(handlers)
}
