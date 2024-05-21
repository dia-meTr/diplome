package postgres

import (
	"context"

	"oss-backend/internal/models"

	"github.com/google/uuid"
)

func (p *Postgres) GetDietByID(ctx context.Context, dietID uuid.UUID) (*models.Diet, error) {
	var diet models.Diet

	err := p.db.NewSelect().
		Model(&diet).
		Where("id = ?", dietID).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return &diet, nil
}

func (p *Postgres) ListDiets(ctx context.Context) ([]*models.Diet, error) {
	var diets []*models.Diet

	err := p.db.NewSelect().
		Model(&diets).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return diets, nil
}

func (p *Postgres) CreateDiet(ctx context.Context, diet *models.Diet) error {
	_, err := p.db.NewInsert().
		Model(diet).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) UpdateDiet(ctx context.Context, diet *models.Diet) error {
	_, err := p.db.NewUpdate().
		Model(diet).
		WherePK().
		Returning("*").
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) DeleteDiet(ctx context.Context, dietID uuid.UUID) error {
	_, err := p.db.NewDelete().
		Model(&models.Diet{}).
		Where("id = ?", dietID).
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}
