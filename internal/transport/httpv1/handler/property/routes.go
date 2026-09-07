package property

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterRoutes(api huma.API, handler *PropertyHandler) {
	huma.Register(api, huma.Operation{
		Method:        http.MethodPost,
		Path:          "/api/v1/properties",
		Summary:       "Создать объект недвижимости",
		Tags:          []string{"Properties"},
		Description:   "Создает новый объект недвижимости для текущего пользователя",
		Security:      []map[string][]string{{"jwt": {}}},
		DefaultStatus: http.StatusCreated,
	}, handler.CreateProperty)

	huma.Register(api, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/api/v1/properties/{id}",
		Summary:     "Получить объект недвижимости",
		Tags:        []string{"Properties"},
		Description: "Возвращает объект недвижимости по ID",
	}, handler.GetProperty)
}
