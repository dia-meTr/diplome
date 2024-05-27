package product

import (
	"context"

	"oss-backend/internal/models"

	"github.com/google/uuid"
)

func (s *Service) GetProductByID(id uuid.UUID) (*models.Product, error) {
	ctx := context.Background()

	return s.repo.Product().GetProductByID(ctx, id)
}

func (s *Service) ListProducts() ([]*models.Product, error) {
	ctx := context.Background()

	return s.repo.Product().ListProducts(ctx)
}

func (s *Service) CreateProduct(product *models.Product) error {
	ctx := context.Background()

	return s.repo.Product().CreateProduct(ctx, product)
}

func (s *Service) UpdateProduct(product *models.Product) error {
	ctx := context.Background()

	return s.repo.Product().UpdateProduct(ctx, product)
}

func (s *Service) DeleteProduct(id uuid.UUID) error {
	ctx := context.Background()

	return s.repo.Product().DeleteProduct(ctx, id)
}
