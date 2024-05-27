package user

import (
	"context"

	"oss-backend/internal/models"

	"github.com/google/uuid"
)

func (s *Service) GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	return s.repo.User().GetUserByID(ctx, userID)
}

func (s *Service) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.repo.User().GetUserByEmail(ctx, email)
}

func (s *Service) Update(ctx context.Context, user *models.User) error {
	return s.repo.User().UpdateUser(ctx, user)
}
