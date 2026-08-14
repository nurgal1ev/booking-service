package httpv1

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

func StartServer() {
	r := chi.NewMux()

	humaCfg := huma.DefaultConfig("Booking Api", "1.0.0")

	api := humachi.New(r, humaCfg)

	huma.Post(api, "/api/v1/auth/register", user.RegisterHandler)
	huma.Post(api, "/api/v1/auth/login", user.LoginHandler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
