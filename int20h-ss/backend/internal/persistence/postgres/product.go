package postgres

import (
	"context"

	"oss-backend/internal/models"

	"github.com/google/uuid"
)

func (p *Postgres) GetProductByID(ctx context.Context, productID uuid.UUID) (*models.Product, error) {
	var product models.Product

	err := p.db.NewSelect().
		Model(&product).
		Where("id = ?", productID).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return &product, nil
}

func (p *Postgres) ListProducts(ctx context.Context) ([]*models.Product, error) {
	var products []*models.Product

	err := p.db.NewSelect().
		Model(&products).
		Scan(ctx)
	if err != nil {
		return nil, p.err(err)
	}

	return products, nil
}

func (p *Postgres) CreateProduct(ctx context.Context, product *models.Product) error {
	_, err := p.db.NewInsert().
		Model(product).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) UpdateProduct(ctx context.Context, product *models.Product) error {
	_, err := p.db.NewUpdate().
		Model(product).
		WherePK().
		Returning("*").
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}

func (p *Postgres) DeleteProduct(ctx context.Context, productID uuid.UUID) error {
	_, err := p.db.NewDelete().
		Model(&models.Product{}).
		Where("id = ?", productID).
		Exec(ctx)
	if err != nil {
		return p.err(err)
	}

	return nil
}
