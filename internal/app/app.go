package app

import (
	"github.com/nurgal1ev/booking-service/internal/config"
	"github.com/nurgal1ev/booking-service/internal/repository/property"
	unitRepo "github.com/nurgal1ev/booking-service/internal/repository/unit"
	userRepo "github.com/nurgal1ev/booking-service/internal/repository/user"
	propertyService "github.com/nurgal1ev/booking-service/internal/service/property"
	unitService "github.com/nurgal1ev/booking-service/internal/service/unit"
	userService "github.com/nurgal1ev/booking-service/internal/service/user"
	"github.com/nurgal1ev/booking-service/internal/transport/httpv1"
	propertyHandler "github.com/nurgal1ev/booking-service/internal/transport/httpv1/handler/property"
	unitHandler "github.com/nurgal1ev/booking-service/internal/transport/httpv1/handler/unit"
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

	unitRepo := unitRepo.NewUnitRepo(db)
	unitService := unitService.NewUnitService(unitRepo, propertyRepository)
	unitHandler := unitHandler.NewUnitHandler(unitService)

	return httpv1.Handlers{
		User:     userHandler,
		Property: propertyHandler,
		Unit:     unitHandler,
	}
}
