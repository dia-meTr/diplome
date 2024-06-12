package tag

import (
	"context"

	"oss-backend/internal/models"
)

func (s *Service) GetTagByID(id int) (*models.Tag, error) {
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

func (s *Service) DeleteTag(id int) error {
	ctx := context.Background()

	return s.repo.Tag().DeleteTag(ctx, id)
}
