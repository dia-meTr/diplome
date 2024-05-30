package user

import (
	"context"
	"fmt"

	"oss-backend/internal/models"

	"github.com/google/uuid"
)

func (s *Service) GetUserByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	resp, err := s.repo.User().GetUserByID(ctx, userID)

	if err != nil {
		return nil, fmt.Errorf("failed executing remote GetUser: %w", err)
	}

	return resp, nil
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	resp, err := s.repo.User().GetUserByEmail(ctx, email)

	if err != nil {
		return nil, fmt.Errorf("failed executing remote GetUser: %w", err)
	}

	return resp, nil
}

func (s *Service) UserExistsByEmail(ctx context.Context, email string) (bool, error) {
	resp, err := s.repo.User().UserExistsByEmail(ctx, email)

	if err != nil {
		return false, fmt.Errorf("failed executing remote GetUser: %w", err)
	}

	return resp, nil
}

func (s *Service) CreateUser(ctx context.Context, user *models.User) error {
	err := s.repo.User().CreateUser(ctx, user)

	return err
}
