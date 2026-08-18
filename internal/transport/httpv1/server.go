package httpv1

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"

	userHandler "github.com/nurgal1ev/booking-service/internal/transport/httpv1/handler/user"
)

type Handlers struct {
	User *userHandler.UserHandler
}

func StartServer(h Handlers) {
	r := chi.NewMux()

	humaCfg := huma.DefaultConfig("Booking Api", "1.0.0")
	humaCfg.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"jwt": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}

	api := humachi.New(r, humaCfg)

	huma.Post(api, "/api/v1/auth/register", h.User.RegisterHandler)
	huma.Post(api, "/api/v1/auth/login", h.User.LoginHandler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
