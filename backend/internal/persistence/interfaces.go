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
}

type Auth interface {
	CreateCredentials(ctx context.Context, credentials *models.UserCredentials) error
}

type User interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	//	ListTeachers(ctx context.Context) ([]models.Teacher, error)
	//	ListStudents(ctx context.Context, opts ...QueryBuilder) ([]models.Student, error)
	//	ListStudentsByGroupID(ctx context.Context, groupID uuid.UUID) ([]models.Student, error)
	//	ListStudentsByFacultyID(ctx context.Context, facultyID uuid.UUID) ([]models.Student, error)

	CreateUser(ctx context.Context, user *models.User) error
	//	CreateTeacher(ctx context.Context, teacher *models.Teacher) error
	//	CreateStudent(ctx context.Context, student *models.Student) error

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
	GetTagByID(ctx context.Context, tagID uuid.UUID) (*models.Tag, error)
	CreateTag(ctx context.Context, tag *models.Tag) error
	UpdateTag(ctx context.Context, tag *models.Tag) error
	DeleteTag(ctx context.Context, tagID uuid.UUID) error
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
