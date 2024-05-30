package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
)

func (s *Service) RefreshToken(ctx context.Context, refreshToken string) (*oauth2.Token, error) {
	parsedToken, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWT.Secret), nil
	})
	if err != nil || !parsedToken.Valid {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fmt.Errorf("token expired: %w", err)
		}

		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims := parsedToken.Claims.(jwt.MapClaims)

	if claims["typ"] != "refresh" {
		return nil, errors.New("you must provide a refresh token")
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		return nil, errors.New("invalid token")
	}

	token, err := s.createBearerToken(ctx, userID, s.cfg.JWT.AccessExpiresIn, s.cfg.JWT.RefreshExpiresIn)
	if err != nil {
		return nil, fmt.Errorf("failed to create bearer token: %w", err)
	}

	return token, nil
}
