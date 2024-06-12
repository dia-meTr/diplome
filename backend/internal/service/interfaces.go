package service

import (
	"context"
	"io"
	"oss-backend/internal/models"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

type Auth interface {
	GetOAuthConsentURL(ctx context.Context) string
	ExchangeProvidersCode(ctx context.Context, code string) (*oauth2.Token, error)
	RefreshToken(ctx context.Context, refreshToken string) (*oauth2.Token, error)
}

type User interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UserExistsByEmail(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, cur *models.User) error
}

type Product interface {
	GetProductByID(id uuid.UUID) (*models.Product, error)
	ListProducts() ([]*models.Product, error)
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uuid.UUID) error
}

type Media interface {
	Upload(ctx context.Context, fileReader io.Reader, extension string) (url string, err error)
}

type Tag interface {
	GetTagByID(id int) (*models.Tag, error)
	ListTags() ([]*models.Tag, error)
	CreateTag(tag *models.Tag) error
	UpdateTag(tag *models.Tag) error
	DeleteTag(id int) error
}

type Dish interface {
	//ListAssignmentsBySubject(ctx context.Context, subjectID uuid.UUID) ([]models.Assignment, error)
	GetDishByID(ctx context.Context, dishID uuid.UUID) (*models.Dish, error)
	GetDishesByTags(ctx context.Context, tagIDs []int) ([]models.Dish, error) // Modified method name to use named parameters
	ListDishes() ([]*models.Dish, error)
	CreateDish(ctx context.Context, dish *models.Dish) error
	UpdateDish(ctx context.Context, dish *models.Dish) error
	DeleteDish(ctx context.Context, dishID uuid.UUID) error
	AddTagToDish(ctx context.Context, dishTag *models.DishTag) error
}

type Order interface {
	GetOrderByID(ctx context.Context, orderID uuid.UUID) (*models.Order, error)
	ListUsersOrders(ctx context.Context, userID uuid.UUID) ([]*models.Order, error)
	ListOrders(ctx context.Context, deliveryDate string, orderStatus string) ([]*models.Order, error)
	CreateOrder(ctx context.Context, order *models.Order) error
	UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, orderStatus string) error
	DeleteOrder(ctx context.Context, orderID uuid.UUID) error
}

type ShoppingCard interface {
	//GetShoppingCardByID(ctx context.Context, cardID uuid.UUID) (*models.ShoppingCard, error)
	GetUsersCardByID(ctx context.Context, userID uuid.UUID) ([]*models.ShoppingCard, error)
	AddItemToCard(ctx context.Context, userID uuid.UUID, dishID uuid.UUID) error
	RemoveItemFromCard(ctx context.Context, userID uuid.UUID, dishID uuid.UUID) error
	UpdateItemAmount(ctx context.Context, userID uuid.UUID, dishID uuid.UUID, amount int) error
	ClearCard(ctx context.Context, userID uuid.UUID) error
	ItemExistsById(ctx context.Context, userID uuid.UUID, dishID uuid.UUID) (int, error)
}
