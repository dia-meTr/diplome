package tag

import (
	"context"

	"oss-backend/internal/models"

	"github.com/google/uuid"
)

func (s *Service) GetTagByID(id uuid.UUID) (*models.Tag, error) {
	ctx := context.Background()

	return s.repo.Tag().GetTagByID(ctx, id)
}

func (s *Service) ListTags() ([]*models.Tag, error) {
	ctx := context.Background()

	return s.repo.Tag().ListTags(ctx)
}

func (s *Service) CreateTag(tag *models.Tag) error {
	ctx := context.Background()

	return s.repo.Tag().CreateTag(ctx, tag)
}

func (s *Service) UpdateTag(tag *models.Tag) error {
	ctx := context.Background()

	return s.repo.Tag().UpdateTag(ctx, tag)
}

func (s *Service) DeleteTag(id uuid.UUID) error {
	ctx := context.Background()

	return s.repo.Tag().DeleteTag(ctx, id)
}
