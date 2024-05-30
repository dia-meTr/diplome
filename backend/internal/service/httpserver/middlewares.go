package httpserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"oss-backend/internal/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

const (
	userIDKey = "user_id"
)

func (s *HTTPServer) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")

		if len(t) != 2 || t[0] != "Bearer" {
			s.respondError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
			return
		}

		authToken := t[1]

		parsedToken, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.cfg.JWT.Secret), nil
		})
		if err != nil || !parsedToken.Valid {
			if errors.Is(err, jwt.ErrTokenExpired) {
				s.respondError(w, http.StatusUnauthorized, fmt.Errorf("token expired"))
				return
			}
			s.respondError(w, http.StatusUnauthorized, fmt.Errorf("Unknown error"))
			return
		}

		claims := parsedToken.Claims.(jwt.MapClaims)

		if claims["typ"] != "access" {
			s.respondError(w, http.StatusUnauthorized, fmt.Errorf("invalid token"))
			return
		}

		userIDStr, ok := claims["sub"].(string)
		if !ok {
			s.respondError(w, http.StatusUnauthorized, fmt.Errorf("invalid token"))
			return
		}

		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			s.respondError(w, http.StatusUnauthorized, fmt.Errorf("invalid token"))
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), userIDKey, userID))

		next.ServeHTTP(w, r)
	})
}

func (s *HTTPServer) roleMiddleware(admin models.Role) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//next.ServeHTTP(w, r)
			//return
			user := r.Context().Value("user").(*models.User)

			if user.UserRole != admin {
				s.respondError(w, http.StatusForbidden, fmt.Errorf("forbidden"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
