package order

import (
	"context"
	"fmt"

	"oss-backend/internal/models"

	"github.com/google/uuid"
)

func (s *Service) GetOrderByID(ctx context.Context, orderID uuid.UUID) (*models.Order, error) {
	order, err := s.repo.Order().GetOrderByID(ctx, orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to get order: %v", err)
	}

	return order, nil
}

func (s *Service) ListUsersOrders(ctx context.Context, userID uuid.UUID) ([]*models.Order, error) {
	orders, err := s.repo.Order().ListUsersOrders(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get order: %v", err)
	}

	return orders, nil
}

func (s *Service) ListOrders(ctx context.Context, deliveryDate string, orderStatus string) ([]*models.Order, error) {

	if deliveryDate == "" && orderStatus == "" {
		orders, err := s.repo.Order().ListOrders(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get order: %v", err)
		}

		return orders, nil
	} else if deliveryDate == "" {
		orders, err := s.repo.Order().ListOrdersByStatus(ctx, orderStatus)
		if err != nil {
			return nil, fmt.Errorf("failed to get order: %v", err)
		}

		return orders, nil
	} else {
		orders, err := s.repo.Order().ListOrdersByDateAndStatus(ctx, deliveryDate, orderStatus)
		if err != nil {
			return nil, fmt.Errorf("failed to get order: %v", err)
		}

		return orders, nil
	}

}

func (s *Service) CreateOrder(ctx context.Context, order *models.Order) error {
	order.OrderStatus = models.NewStatus

	if err := s.repo.Order().CreateOrder(ctx, order); err != nil {
		return fmt.Errorf("failed to create order: %v", err)
	}

	cart, err := s.repo.ShoppingCart().GetUsersCartByID(ctx, order.UserID)
	if err != nil {
		return fmt.Errorf("failed to create order: %v", err)
	}

	orderDishes := make([]models.OrderDish, len(cart))

	for i, dish := range cart {
		orderDishes[i] = models.OrderDish{
			OrderID: order.ID,
			DishID:  dish.DishID,
			Amount:  dish.Amount,
		}
	}

	if err := s.repo.Order().AddDishesToOrder(ctx, orderDishes); err != nil {
		return fmt.Errorf("failed to create order: %v", err)
	}

	orderID := order.ID
	fmt.Println(orderID)

	return nil
}

func (s *Service) UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, orderStatus string) error {
	if err := s.repo.Order().UpdateOrderStatus(ctx, orderID, orderStatus); err != nil {
		return fmt.Errorf("failed to update order: %v", err)
	}

	return nil
}

func (s *Service) DeleteOrder(ctx context.Context, orderID uuid.UUID) error {
	if err := s.repo.Order().DeleteOrder(ctx, orderID); err != nil {
		return fmt.Errorf("failed to delete order: %v", err)
	}

	return nil
}
