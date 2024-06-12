package persistence

import (
	"context"

	"oss-backend/internal/models"

	"github.com/google/uuid"
)

type Repo interface {
	Auth() Auth
	User() User
	Product() Product
	Tag() Tag
	Dish() Dish
	Order() Order
	ShoppingCart() ShoppingCard
}

type Auth interface {
	CreateCredentials(ctx context.Context, credentials *models.UserCredentials) error
}

type User interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UserExistsByEmail(ctx context.Context, email string) (bool, error)
	UpdateUser(ctx context.Context, user *models.User) error
}

type Product interface {
	ListProducts(ctx context.Context) ([]*models.Product, error)
	GetProductByID(ctx context.Context, productID uuid.UUID) (*models.Product, error)
	CreateProduct(ctx context.Context, product *models.Product) error
	UpdateProduct(ctx context.Context, product *models.Product) error
	DeleteProduct(ctx context.Context, productID uuid.UUID) error
}

type Tag interface {
	ListTags(ctx context.Context) ([]*models.Tag, error)
	GetTagByID(ctx context.Context, tagID int) (*models.Tag, error)
	CreateTag(ctx context.Context, tag *models.Tag) error
	UpdateTag(ctx context.Context, tag *models.Tag) error
	DeleteTag(ctx context.Context, tagID int) error
}

type Dish interface {
	GetDishByID(ctx context.Context, dishID uuid.UUID) (*models.Dish, error)
	ListDishes(ctx context.Context) ([]*models.Dish, error)
	CreateDish(ctx context.Context, dish *models.Dish) error
	UpdateDish(ctx context.Context, dish *models.Dish) error
	DeleteDish(ctx context.Context, dishID uuid.UUID) error
	//ListDishesBySubject(ctx context.Context, subjectID uuid.UUID) ([]models.Dish, error)
	AddTagToDish(ctx context.Context, dishTag *models.DishTag) error
	GetDishesByTags(ctx context.Context, tagIDs []int) ([]models.Dish, error)
}

type Order interface {
	GetOrderByID(ctx context.Context, orderID uuid.UUID) (*models.Order, error)
	ListOrders(ctx context.Context) ([]*models.Order, error)
	ListUsersOrders(ctx context.Context, userID uuid.UUID) ([]*models.Order, error)
	ListOrdersByDateAndStatus(ctx context.Context, deliveryDate string, orderStatus string) ([]*models.Order, error)
	ListOrdersByStatus(ctx context.Context, orderStatus string) ([]*models.Order, error)
	CreateOrder(ctx context.Context, order *models.Order) error
	UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, orderStatus string) error
	DeleteOrder(ctx context.Context, orderID uuid.UUID) error
	AddDishesToOrder(ctx context.Context, orderDishes []models.OrderDish) error
}

type ShoppingCard interface {
	GetUsersCartByID(ctx context.Context, userID uuid.UUID) ([]*models.ShoppingCard, error)
	AddItemToCard(ctx context.Context, cardItem *models.ShoppingCard) error
	RemoveItemFromCard(ctx context.Context, userID uuid.UUID, dishID uuid.UUID) error
	UpdateItemAmount(ctx context.Context, userID uuid.UUID, dishID uuid.UUID, amount int) error
	ClearCard(ctx context.Context, userID uuid.UUID) error
	ItemExistsById(ctx context.Context, userID uuid.UUID, dishID uuid.UUID) (int, error)
}
