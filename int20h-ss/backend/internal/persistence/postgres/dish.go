package postgres

import (
	"context"

	"oss-backend/internal/models"

	"github.com/google/uuid"
)

func (p *Postgres) GetDishByID(ctx context.Context, id uuid.UUID) (*models.Dish, error) {
	var dish models.Dish

	if err := p.db.NewSelect().
		Model(&dish).
		Where("id = ?", id).
		Scan(ctx); err != nil {
		return nil, p.err(err)
	}
	return &dish, nil
}

func (p *Postgres) CreateDish(ctx context.Context, dish *models.Dish) error {
	if _, err := p.db.NewInsert().
		Model(dish).
		Returning("*").
		Exec(ctx); err != nil {
		return p.err(err)
	}
	return nil
}

func (p *Postgres) UpdateDish(ctx context.Context, dish *models.Dish) error {
	if _, err := p.db.NewUpdate().
		Model(dish).
		Returning("*").
		WherePK().
		Exec(ctx); err != nil {
		return p.err(err)
	}
	return nil
}

func (p *Postgres) DeleteDish(ctx context.Context, id uuid.UUID) error {
	if _, err := p.db.NewDelete().
		Model(&models.Dish{}).
		Where("id = ?", id).
		Exec(ctx); err != nil {
		return p.err(err)
	}
	return nil
}

func (p *Postgres) AddTagToDish(ctx context.Context, dishTag *models.DishTag) error {
	if _, err := p.db.NewInsert().
		Model(dishTag).
		Returning("*").
		Exec(ctx); err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) AddDietToDish(ctx context.Context, dishDiet *models.DishDiet) error {
	if _, err := p.db.NewInsert().
		Model(dishDiet).
		Returning("*").
		Exec(ctx); err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) ListDishesByCategory(ctx context.Context, categoryID uuid.UUID) ([]models.Dish, error) {
	var dishes []models.Dish

	err := p.db.NewSelect().
		Model(&dishes).
		Where("category_id = ?", categoryID).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return dishes, nil
}
