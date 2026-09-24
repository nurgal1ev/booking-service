package unit

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/nurgal1ev/booking-service/internal/transport/middleware"
)

func RegisterRoutes(api huma.API, handler *UnitHandler, authSecret string) {
	authMiddleware := huma.Middlewares{
		middleware.AuthMiddleware(api, authSecret),
	}

	huma.Register(api, huma.Operation{
		Method:        http.MethodPost,
		Path:          "/api/v1/properties/{id}/units",
		Summary:       "Создать комнату",
		Tags:          []string{"Units"},
		Description:   "Создает новую комнату (номер, домик) для указанного объекта недвижимости. Доступно только владельцу объекта.",
		Security:      []map[string][]string{{"jwt": {}}},
		Middlewares:   authMiddleware,
		DefaultStatus: http.StatusCreated,
	}, handler.CreateUnitHandler)
}
