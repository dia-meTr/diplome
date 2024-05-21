package persistence

import (
	"context"

	"oss-backend/internal/models"

	"github.com/google/uuid"
)

type Repo interface {
	Auth() Auth
	User() User
	Faculty() Faculty
	Group() Group
	Subject() Subject
	Event() Event
	Notification() Notification
	Activity() Activity
	Assignment() Assignment
	Product() Product
	Diet() Diet
	Tag() Tag
	Dish() Dish
}

type Auth interface {
	CreateCredentials(ctx context.Context, credentials *models.UserCredentials) error
}

type User interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	ListTeachers(ctx context.Context) ([]models.Teacher, error)
	ListStudents(ctx context.Context, opts ...QueryBuilder) ([]models.Student, error)
	ListStudentsByGroupID(ctx context.Context, groupID uuid.UUID) ([]models.Student, error)
	ListStudentsByFacultyID(ctx context.Context, facultyID uuid.UUID) ([]models.Student, error)

	CreateUser(ctx context.Context, user *models.User) error
	CreateTeacher(ctx context.Context, teacher *models.Teacher) error
	CreateStudent(ctx context.Context, student *models.Student) error

	UpdateUser(ctx context.Context, user *models.User) error
}

type Group interface {
	ListGroups(ctx context.Context) ([]*models.Group, error)
	ListGroupsByFacultyID(ctx context.Context, facultyID uuid.UUID) ([]*models.Group, error)
	UpdateGroup(ctx context.Context, group *models.Group) error
	CreateGroup(ctx context.Context, group *models.Group) error
	DeleteGroup(ctx context.Context, groupID uuid.UUID) error
}

type Faculty interface {
	ListFaculties(ctx context.Context) ([]*models.Faculty, error)
	GetFacultyByID(ctx context.Context, facultyID uuid.UUID) (*models.Faculty, error)
	CreateFaculty(ctx context.Context, faculty *models.Faculty) error
	UpdateFaculty(ctx context.Context, faculty *models.Faculty) error
	DeleteFaculty(ctx context.Context, facultyID uuid.UUID) error
}

type Subject interface {
	ListSubjects(ctx context.Context, opts *models.SubjectFilterDTO) ([]*models.Subject, error)
	CreateSubject(ctx context.Context, subject *models.Subject) error
	UpdateSubject(ctx context.Context, subject *models.Subject) error
	DeleteSubject(ctx context.Context, subjectID uuid.UUID) error
}

type Event interface {
	ListEvents(ctx context.Context, opts ...QueryBuilder) ([]*models.Event, error)
	CreateEvent(ctx context.Context, event *models.Event) error
	UpdateEvent(ctx context.Context, event *models.Event) error
	DeleteEvent(ctx context.Context, eventID uuid.UUID) error
}

type Notification interface {
	GetTemplates(ctx context.Context) ([]models.NotificationTemplate, error)
	CreateTemplate(ctx context.Context, template *models.NotificationTemplate) error
	UpdateTemplate(ctx context.Context, template *models.NotificationTemplate) error
	DeleteTemplate(ctx context.Context, templateID uuid.UUID) error
}

type Activity interface {
	GetActivityByID(ctx context.Context, id uuid.UUID) (*models.Activity, error)
	ListActivitiesByStudent(ctx context.Context, id uuid.UUID) ([]models.Activity, error)
	CreateActivity(ctx context.Context, activity *models.Activity) error
	UpdateActivity(ctx context.Context, activity *models.Activity) error
	DeleteActivity(ctx context.Context, activityID uuid.UUID) error
}

type Assignment interface {
	GetAssignmentByID(ctx context.Context, assignmentID uuid.UUID) (*models.Assignment, error)
	ListAssignmentsBySubject(ctx context.Context, subjectID uuid.UUID) ([]models.Assignment, error)
	CreateAssignment(ctx context.Context, assignment *models.Assignment) error
	CreateSubmittedAssignment(ctx context.Context, submittedAssigment *models.SubmittedAssigment) error
	UpdateAssignment(ctx context.Context, assignment *models.Assignment) error
	DeleteAssignment(ctx context.Context, assignmentID uuid.UUID) error
}

type Product interface {
	ListProducts(ctx context.Context) ([]*models.Product, error)
	GetProductByID(ctx context.Context, productID uuid.UUID) (*models.Product, error)
	CreateProduct(ctx context.Context, product *models.Product) error
	UpdateProduct(ctx context.Context, product *models.Product) error
	DeleteProduct(ctx context.Context, productID uuid.UUID) error
}

type Diet interface {
	ListDiets(ctx context.Context) ([]*models.Diet, error)
	GetDietByID(ctx context.Context, dietID uuid.UUID) (*models.Diet, error)
	CreateDiet(ctx context.Context, diet *models.Diet) error
	UpdateDiet(ctx context.Context, diet *models.Diet) error
	DeleteDiet(ctx context.Context, dietID uuid.UUID) error
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
	CreateDish(ctx context.Context, dish *models.Dish) error
	UpdateDish(ctx context.Context, dish *models.Dish) error
	DeleteDish(ctx context.Context, dishID uuid.UUID) error
	//ListDishesBySubject(ctx context.Context, subjectID uuid.UUID) ([]models.Dish, error)
	AddTagToDish(ctx context.Context, dishTag *models.DishTag) error
	AddDietToDish(ctx context.Context, dishDiet *models.DishDiet) error
}
