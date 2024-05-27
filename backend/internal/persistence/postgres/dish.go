package postgres

import (
	"context"

	"oss-backend/internal/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (p *Postgres) GetDishByID(ctx context.Context, id uuid.UUID) (*models.Dish, error) {
	var dish models.Dish

	if err := p.db.NewSelect().
		Model(&dish).
		Scan(ctx); err != nil {
		return nil, p.err(err)
	}
	return &dish, nil
}

func (p *Postgres) GetDishesByTags(ctx context.Context, tagIDs []int) ([]models.Dish, error) {
	var dishes []models.Dish

	err := p.db.NewSelect().
		Model(&dishes).
		Relation("Tags", func(q *bun.SelectQuery) *bun.SelectQuery {
			q = q.Where("tag.id IN (?)", bun.In(tagIDs))
			//				Group("dish_tag.dish_id").
			//				Having("COUNT(DISTINCT tag.id) = 2")
			return q
		}).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return dishes, nil
}

func (p *Postgres) ListDishes(ctx context.Context) ([]*models.Dish, error) {
	var dishes []*models.Dish

	err := p.db.NewSelect().
		Model(&dishes).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return dishes, nil
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
