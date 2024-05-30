package postgres

import (
	"context"
	"fmt"

	"oss-backend/internal/models"

	"github.com/google/uuid"
)

func (p *Postgres) GetUsersCartByID(ctx context.Context, userID uuid.UUID) ([]*models.ShoppingCard, error) {
	var card []*models.ShoppingCard

	err := p.db.NewSelect().
		Model(&card).
		Where("user_id = ?", userID).
		Relation("Dish").
		//, func(q *bun.SelectQuery) *bun.SelectQuery {
		//q = q.Where("tag.id IN (?)", bun.In(tagIDs))
		//				Group("dish_tag.dish_id").
		//				Having("COUNT(DISTINCT tag.id) = 2")
		//return q
		//}
		//).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return card, nil
}

func (p *Postgres) AddItemToCard(ctx context.Context, cardItem *models.ShoppingCard) error {
	_, err := p.db.NewInsert().
		Model(cardItem).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) RemoveItemFromCard(ctx context.Context, userID uuid.UUID, dishID uuid.UUID) error {
	_, err := p.db.NewDelete().
		Model(&models.ShoppingCard{}).
		Where("user_id = ?", userID).
		Where("dish_id = ?", dishID).
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) UpdateItemAmount(ctx context.Context, userID uuid.UUID, dishID uuid.UUID, amount int) error {
	_, err := p.db.NewUpdate().
		Model(&models.ShoppingCard{}).
		Set("amount = ?", amount).
		Where("user_id = ?", userID).
		Where("dish_id = ?", dishID).
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}
	return nil
}

func (p *Postgres) ClearCard(ctx context.Context, userID uuid.UUID) error {
	_, err := p.db.NewDelete().
		Model(&models.ShoppingCard{}).
		Where("user_id = ?", userID).
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) ItemExistsById(ctx context.Context, userID uuid.UUID, dishID uuid.UUID) (int, error) {
	var item models.ShoppingCard

	err := p.db.NewSelect().
		Model(&item).
		Where("user_id = ?", userID).
		Where("dish_id = ?", dishID).
		Scan(ctx)

	if err != nil {
		return 0, p.err(err)
	}

	fmt.Print(item.Amount)

	return item.Amount, nil
}
