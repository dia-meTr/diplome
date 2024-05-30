package dish

import (
	"context"
	"fmt"

	"oss-backend/internal/models"

	"github.com/google/uuid"
)

func (s *Service) GetDishByID(ctx context.Context, dishID uuid.UUID) (*models.Dish, error) {
	dish, err := s.repo.Dish().GetDishByID(ctx, dishID)
	if err != nil {
		return nil, fmt.Errorf("failed to get dish: %v", err)
	}

	return dish, nil
}

func (s *Service) GetDishesByTags(ctx context.Context, tagIDs []int) ([]models.Dish, error) {
	dishes, err := s.repo.Dish().GetDishesByTags(ctx, tagIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to get dishes by tags: %v", err)
	}

	return dishes, nil
}

func (s *Service) ListDishes() ([]*models.Dish, error) {
	ctx := context.Background()

	return s.repo.Dish().ListDishes(ctx)
}

func (s *Service) CreateDish(ctx context.Context, dish *models.Dish) error {

	if err := s.repo.Dish().CreateDish(ctx, dish); err != nil {
		return fmt.Errorf("failed to create dish: %v", err)
	}

	return nil
}

func (s *Service) AddTagToDish(ctx context.Context, dishTag *models.DishTag) error {

	if err := s.repo.Dish().AddTagToDish(ctx, dishTag); err != nil {
		return fmt.Errorf("failed to add tag to dish: %v", err)
	}

	return nil

}

func (s *Service) UpdateDish(ctx context.Context, dish *models.Dish) error {
	if err := s.repo.Dish().UpdateDish(ctx, dish); err != nil {
		return fmt.Errorf("failed to update dish: %v", err)
	}

	return nil
}

func (s *Service) DeleteDish(ctx context.Context, dishID uuid.UUID) error {
	if err := s.repo.Dish().DeleteDish(ctx, dishID); err != nil {
		return fmt.Errorf("failed to delete dish: %v", err)
	}

	return nil
}
