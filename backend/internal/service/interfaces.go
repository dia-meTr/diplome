package service

import (
	"context"
	"io"
	"oss-backend/internal/models"

	"github.com/google/uuid"
)

type Auth interface {
	GenerateToken(ctx context.Context, user *models.User) (*models.UserCredentials, error)
	Login(ctx context.Context, accessToken string) (*models.User, error)
}

type User interface {
	GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	//	ListTeachers(ctx context.Context) ([]models.Teacher, error)
	//	ListStudents(ctx context.Context, filter models.FilterUserDTO) ([]models.Student, error)
	//	ListStudentsByGroupID(ctx context.Context, groupID uuid.UUID) ([]models.Student, error)
	//	InviteTeacher(ctx context.Context, dto *models.InviteTeacherDTO) error
	//	InviteStudent(ctx context.Context, dto *models.InviteStudentDTO) error
	Update(ctx context.Context, user *models.User) error
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
	GetTagByID(id uuid.UUID) (*models.Tag, error)
	ListTags() ([]*models.Tag, error)
	CreateTag(tag *models.Tag) error
	UpdateTag(tag *models.Tag) error
	DeleteTag(id uuid.UUID) error
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
