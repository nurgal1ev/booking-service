package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIDKey contextKey = "user_id"
const UserRoleKey contextKey = "user_role"

func AuthMiddleware(api huma.API, secret string) func(huma.Context, func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		authHeader := ctx.Header("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Missing or invalid token")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf(
					"unexpected signing method: %v",
					token.Header["alg"],
				)
			}

			return []byte(secret), nil
		})

		if err != nil {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Invalid token")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Invalid token")
			return
		}

		userID, ok := claims["user_id"].(float64)
		if !ok {
			huma.WriteErr(api, ctx, http.StatusUnauthorized, "Invalid token claims")
			return
		}

		role, _ := claims["role"].(string)

		ctx = huma.WithValue(ctx, UserIDKey, uint(userID))
		ctx = huma.WithValue(ctx, UserRoleKey, role)

		next(ctx)
	}
}

func GetUserIDFromContext(ctx context.Context) (uint, bool) {
	userID, ok := ctx.Value(UserIDKey).(uint)
	return userID, ok
}

func GetUserRoleFromContext(ctx context.Context) (string, bool) {
	role, ok := ctx.Value(UserRoleKey).(string)
	return role, ok
}
