package diet

import (
	"context"

	"oss-backend/internal/models"

	"github.com/google/uuid"
)

func (s *Service) GetDietByID(id uuid.UUID) (*models.Diet, error) {
	ctx := context.Background()

	return s.repo.Diet().GetDietByID(ctx, id)
}

func (s *Service) ListDiets() ([]*models.Diet, error) {
	ctx := context.Background()

	return s.repo.Diet().ListDiets(ctx)
}

func (s *Service) CreateDiet(diet *models.Diet) error {
	ctx := context.Background()

	return s.repo.Diet().CreateDiet(ctx, diet)
}

func (s *Service) UpdateDiet(diet *models.Diet) error {
	ctx := context.Background()

	return s.repo.Diet().UpdateDiet(ctx, diet)
}

func (s *Service) DeleteDiet(id uuid.UUID) error {
	ctx := context.Background()

	return s.repo.Diet().DeleteDiet(ctx, id)
}
