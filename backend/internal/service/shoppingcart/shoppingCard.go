package shoppingcart

import (
	"context"
	"fmt"

	"oss-backend/internal/models"

	"github.com/google/uuid"
)

func (s *Service) GetUsersCardByID(ctx context.Context, userID uuid.UUID) ([]*models.ShoppingCard, error) {
	card, err := s.repo.ShoppingCart().GetUsersCartByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get card: %v", err)
	}

	return card, nil
}

func (s *Service) AddItemToCard(ctx context.Context, userID uuid.UUID, dishID uuid.UUID) error {

	amount, err := s.repo.ShoppingCart().ItemExistsById(ctx, userID, dishID)
	if err != nil {
		return fmt.Errorf("failed to check if user exists: %w", err)
	}
	fmt.Print(amount)

	if amount == 0 {

		item := &models.ShoppingCard{
			UserID: userID,
			DishID: dishID,
			Amount: 1,
		}

		if err := s.repo.ShoppingCart().AddItemToCard(ctx, item); err != nil {
			return fmt.Errorf("failed to create dish: %v", err)
		}
	} else {
		if err := s.repo.ShoppingCart().UpdateItemAmount(ctx, userID, dishID, amount+1); err != nil {
			return fmt.Errorf("failed to update dish: %v", err)
		}
	}

	return nil
}

func (s *Service) ItemExistsById(ctx context.Context, userID uuid.UUID, dishID uuid.UUID) (int, error) {

	amount, err := s.repo.ShoppingCart().ItemExistsById(ctx, userID, dishID)
	if err != nil {
		return 0, fmt.Errorf("failed to check if user exists: %w", err)
	}
	fmt.Print(amount)

	return amount, nil
}

func (s *Service) UpdateItemAmount(ctx context.Context, userID uuid.UUID, dishID uuid.UUID, amount int) error {
	if err := s.repo.ShoppingCart().UpdateItemAmount(ctx, userID, dishID, amount); err != nil {
		return fmt.Errorf("failed to update dish: %v", err)
	}

	return nil
}

func (s *Service) RemoveItemFromCard(ctx context.Context, userID uuid.UUID, dishID uuid.UUID) error {
	if err := s.repo.ShoppingCart().RemoveItemFromCard(ctx, userID, dishID); err != nil {
		return fmt.Errorf("failed to delete dish: %v", err)
	}

	return nil
}

func (s *Service) ClearCard(ctx context.Context, userID uuid.UUID) error {
	if err := s.repo.ShoppingCart().ClearCard(ctx, userID); err != nil {
		return fmt.Errorf("failed to clear card: %v", err)
	}

	return nil
}
