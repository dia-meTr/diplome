package postgres

import (
	"context"

	"oss-backend/internal/models"

	"github.com/google/uuid"
)

func (p *Postgres) GetTagByID(ctx context.Context, tagID uuid.UUID) (*models.Tag, error) {
	var tag models.Tag

	err := p.db.NewSelect().
		Model(&tag).
		Where("id = ?", tagID).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return &tag, nil
}

func (p *Postgres) ListTags(ctx context.Context) ([]*models.Tag, error) {
	var tags []*models.Tag

	err := p.db.NewSelect().
		Model(&tags).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return tags, nil
}

func (p *Postgres) CreateTag(ctx context.Context, tag *models.Tag) error {
	_, err := p.db.NewInsert().
		Model(tag).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) UpdateTag(ctx context.Context, tag *models.Tag) error {
	_, err := p.db.NewUpdate().
		Model(tag).
		WherePK().
		Returning("*").
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) DeleteTag(ctx context.Context, tagID uuid.UUID) error {
	_, err := p.db.NewDelete().
		Model(&models.Tag{}).
		Where("id = ?", tagID).
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}
