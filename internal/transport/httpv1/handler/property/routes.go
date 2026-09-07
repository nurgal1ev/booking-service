package property

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/nurgal1ev/booking-service/internal/transport/middleware"
)

func RegisterRoutes(api huma.API, handler *PropertyHandler, authSecret string) {
	authMiddleware := huma.Middlewares{
		middleware.AuthMiddleware(api, authSecret),
	}

	huma.Register(api, huma.Operation{
		Method:        http.MethodPost,
		Path:          "/api/v1/properties",
		Summary:       "Создать объект недвижимости",
		Tags:          []string{"Properties"},
		Description:   "Создает новый объект недвижимости для текущего пользователя",
		Security:      []map[string][]string{{"jwt": {}}},
		DefaultStatus: http.StatusCreated,
		Middlewares:   authMiddleware,
	}, handler.CreateProperty)

	huma.Register(api, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/api/v1/properties/{id}",
		Summary:     "Получить объект недвижимости",
		Tags:        []string{"Properties"},
		Description: "Возвращает объект недвижимости по ID",
	}, handler.GetProperty)
}
