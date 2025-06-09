package middleware

import (
	"context"
	"net/http"
	"strings"

	"leaguies_backend/utils"
)

type contextKey string

const userCtxKey contextKey = "userID"

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		tokenStr := parts[1]

		claims, err := utils.ParseToken(tokenStr)

		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Add userID to the request context
		ctx := context.WithValue(r.Context(), userCtxKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserID extracts the user ID from the request context
func GetUserID(r *http.Request) (uint, bool) {
	userID, ok := r.Context().Value(userCtxKey).(uint)
	return userID, ok
}
