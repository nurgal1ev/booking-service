package app

import (
	"github.com/nurgal1ev/booking-service/internal/config"
	"github.com/nurgal1ev/booking-service/internal/repository/property"
	userRepo "github.com/nurgal1ev/booking-service/internal/repository/user"
	propertyService "github.com/nurgal1ev/booking-service/internal/service/property"
	userService "github.com/nurgal1ev/booking-service/internal/service/user"
	"github.com/nurgal1ev/booking-service/internal/transport/httpv1"
	propertyHandler "github.com/nurgal1ev/booking-service/internal/transport/httpv1/handler/property"
	userHandler "github.com/nurgal1ev/booking-service/internal/transport/httpv1/handler/user"
	"gorm.io/gorm"
)

func InitHandlers(db *gorm.DB, cfg *config.Config) httpv1.Handlers {
	userRepository := userRepo.NewUserRepo(db)
	userService := userService.NewUserService(userRepository, cfg.Auth.Secret)
	userHandler := userHandler.NewUserHandler(userService)

	propertyRepository := property.NewPropertyRepo(db)
	propertyService := propertyService.NewPropertyService(propertyRepository, userRepository)
	propertyHandler := propertyHandler.NewPropertyHandler(propertyService)

	return httpv1.Handlers{
		User:     userHandler,
		Property: propertyHandler,
	}
}
