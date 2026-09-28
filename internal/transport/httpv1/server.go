package httpv1

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/nurgal1ev/booking-service/internal/config"
	propertyHandler "github.com/nurgal1ev/booking-service/internal/transport/httpv1/handler/property"

	userHandler "github.com/nurgal1ev/booking-service/internal/transport/httpv1/handler/user"

	unitHandler "github.com/nurgal1ev/booking-service/internal/transport/httpv1/handler/unit"
)

type Handlers struct {
	User     *userHandler.UserHandler
	Property *propertyHandler.PropertyHandler
	Unit     *unitHandler.UnitHandler
}

func StartServer(h Handlers, cfg *config.Config) {
	r := chi.NewMux()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	humaCfg := huma.DefaultConfig("Booking Api", "1.0.0")
	humaCfg.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"jwt": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}

	api := humachi.New(r, humaCfg)
	userHandler.RegisterRoutes(api, h.User)
	propertyHandler.RegisterRoutes(api, h.Property, cfg.Auth.Secret)

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
