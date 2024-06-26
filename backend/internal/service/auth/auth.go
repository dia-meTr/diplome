package auth

import (
	"context"
	"fmt"
	"time"

	"oss-backend/internal/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"
)

func (s *Service) GetOAuthConsentURL(ctx context.Context) string {
	return s.googleOAuthCfg.AuthCodeURL(randString(64))
}

func (s *Service) ExchangeProvidersCode(ctx context.Context, code string) (*oauth2.Token, error) {
	googleToken, err := s.googleOAuthCfg.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange token: %w", err)
	}

	// Parse ID token
	idTokenPayload, err := idtoken.ParsePayload(googleToken.Extra("id_token").(string))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ID token: %w", err)
	}

	email := idTokenPayload.Claims["email"].(string)

	exists, err := s.userSrv.UserExistsByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to check if user exists: %w", err)
	}

	var user *models.User

	if exists {
		user, err = s.userSrv.GetUserByEmail(ctx, email)
		if err != nil {
			return nil, fmt.Errorf("failed to get user by email: %w", err)
		}
	} else {
		lastName, _ := idTokenPayload.Claims["family_name"].(string)

		user = &models.User{
			Email:     email,
			FirstName: idTokenPayload.Claims["given_name"].(string),
			LastName:  lastName,
			UserRole:  "customer",
		}

		err = s.userSrv.CreateUser(ctx, user)
		if err != nil {
			return nil, fmt.Errorf("failed to create user: %w", err)
		}
	}

	token, err := s.createBearerToken(ctx, user.ID.String(), s.cfg.JWT.AccessExpiresIn, s.cfg.JWT.RefreshExpiresIn)
	if err != nil {
		return nil, fmt.Errorf("failed to own oauth token: %w", err)
	}

	return token, nil
}

func (s *Service) createBearerToken(ctx context.Context, userID string, accessExpiresIn, refreshExpiresIn int) (*oauth2.Token, error) {
	accessToken, err := s.createAccessToken(userID, time.Duration(accessExpiresIn)*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("failed to create JWT: %w", err)
	}

	refreshToken, err := s.createRefreshToken(ctx, userID, time.Duration(refreshExpiresIn)*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("failed to create JWT: %w", err)
	}

	return &oauth2.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		Expiry:       time.Now().Add(time.Duration(accessExpiresIn) * time.Hour),
	}, nil
}

func (s *Service) createAccessToken(userID string, expiresIn time.Duration) (string, error) {
	now := time.Now()

	claims := jwt.MapClaims{
		"sub": userID,
		"iat": now.Unix(),
		"exp": now.Add(expiresIn).Unix(),
		"typ": "access",
	}

	return s.createAndSignToken(claims)
}

func (s *Service) createRefreshToken(ctx context.Context, userID string, expiresIn time.Duration) (string, error) {
	now := time.Now()

	claims := jwt.MapClaims{
		"sub": userID,
		"iat": now.Unix(),
		"exp": now.Add(expiresIn).Unix(),
		"typ": "refresh",
	}

	return s.createAndSignToken(claims)
}

func (s *Service) createAndSignToken(claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := token.SignedString([]byte(s.cfg.JWT.Secret))
	if err != nil {
		return "", fmt.Errorf("failed to sign: %w", err)
	}

	return signedString, nil

}
