package postgres

import (
	"context"

	"oss-backend/internal/models"

	"github.com/google/uuid"
)

func (p *Postgres) GetOrderByID(ctx context.Context, orderID uuid.UUID) (*models.Order, error) {
	var order models.Order

	err := p.db.NewSelect().
		Model(&order).
		Where("id = ?", orderID).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return &order, nil
}

func (p *Postgres) ListOrders(ctx context.Context) ([]*models.Order, error) {
	var orders []*models.Order

	err := p.db.NewSelect().
		Model(&orders).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return orders, nil
}

func (p *Postgres) ListUsersOrders(ctx context.Context, userID uuid.UUID) ([]*models.Order, error) {
	var orders []*models.Order

	err := p.db.NewSelect().
		Model(&orders).
		Where("user_id = ?", userID).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return orders, nil
}

func (p *Postgres) ListOrdersByDateAndStatus(ctx context.Context, deliveryDate string, orderStatus string) ([]*models.Order, error) {
	var orders []*models.Order

	err := p.db.NewSelect().
		Model(&orders).
		Where("delivery_date < ?", deliveryDate).
		Where("order_status = ?", orderStatus).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return orders, nil
}

func (p *Postgres) ListOrdersByStatus(ctx context.Context, orderStatus string) ([]*models.Order, error) {
	var orders []*models.Order

	err := p.db.NewSelect().
		Model(&orders).
		Where("order_status = ?", orderStatus).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return orders, nil
}

func (p *Postgres) CreateOrder(ctx context.Context, order *models.Order) error {
	_, err := p.db.NewInsert().
		Model(order).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, orderStatus string) error {
	_, err := p.db.NewUpdate().
		Model(&models.Order{}).
		Set("order_status", orderStatus).
		Where("id = ?", orderID).
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}
	return nil
}

func (p *Postgres) DeleteOrder(ctx context.Context, orderID uuid.UUID) error {
	_, err := p.db.NewDelete().
		Model(&models.Order{}).
		Where("id = ?", orderID).
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) AddDishesToOrder(ctx context.Context, orderDishes []models.OrderDish) error {
	_, err := p.db.NewInsert().
		Model(&orderDishes).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}
