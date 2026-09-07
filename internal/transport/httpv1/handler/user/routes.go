package user

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterRoutes(api huma.API, handler *UserHandler) {
	huma.Register(api, huma.Operation{
		Method:        http.MethodPost,
		Path:          "/api/v1/auth/register",
		Summary:       "Регистрация",
		Tags:          []string{"Users"},
		DefaultStatus: http.StatusCreated,
	}, handler.RegisterHandler)

	huma.Register(api, huma.Operation{
		Method:  http.MethodPost,
		Path:    "/api/v1/auth/login",
		Summary: "Авторизация",
		Tags:    []string{"Users"},
	}, handler.LoginHandler)
}
