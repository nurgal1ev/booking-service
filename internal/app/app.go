package app

import (
	"github.com/nurgal1ev/booking-service/internal/config"
	userRepo "github.com/nurgal1ev/booking-service/internal/repository/user"
	userService "github.com/nurgal1ev/booking-service/internal/service/user"
	"github.com/nurgal1ev/booking-service/internal/transport/httpv1"
	userHandler "github.com/nurgal1ev/booking-service/internal/transport/httpv1/handler/user"
	"gorm.io/gorm"
)

func InitHandlers(db *gorm.DB, cfg *config.Config) httpv1.Handlers {
	userRepository := userRepo.NewUserRepo(db)
	userService := userService.NewUserService(userRepository, cfg.Auth.Secret)
	userHandler := userHandler.NewUserHandler(userService)

	return httpv1.Handlers{
		User: userHandler,
	}
}
