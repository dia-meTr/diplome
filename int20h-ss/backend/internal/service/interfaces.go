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
	ListTeachers(ctx context.Context) ([]models.Teacher, error)
	ListStudents(ctx context.Context, filter models.FilterUserDTO) ([]models.Student, error)
	ListStudentsByGroupID(ctx context.Context, groupID uuid.UUID) ([]models.Student, error)
	InviteTeacher(ctx context.Context, dto *models.InviteTeacherDTO) error
	InviteStudent(ctx context.Context, dto *models.InviteStudentDTO) error
	Update(ctx context.Context, user *models.User) error
}

type Group interface {
	ListGroups(ctx context.Context) ([]*models.Group, error)
	ListGroupsByFacultyID(ctx context.Context, facultyID uuid.UUID) ([]*models.Group, error)
	UpdateGroup(ctx context.Context, group *models.Group) error
	CreateGroup(ctx context.Context, group *models.Group) error
	DeleteGroup(ctx context.Context, groupID uuid.UUID) error
}

type Faculty interface {
	GetFacultyByID(id uuid.UUID) (*models.Faculty, error)
	ListFaculties() ([]*models.Faculty, error)
	CreateFaculty(faculty *models.Faculty) error
	UpdateFaculty(faculty *models.Faculty) error
	DeleteFaculty(id uuid.UUID) error
}

type Product interface {
	GetProductByID(id uuid.UUID) (*models.Product, error)
	ListProducts() ([]*models.Product, error)
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uuid.UUID) error
}

type Subject interface {
	ListSubjects(ctx context.Context, subject *models.SubjectFilterDTO) ([]*models.Subject, error)
	CreateSubject(ctx context.Context, subject *models.Subject) error
	UpdateSubject(ctx context.Context, subject *models.Subject) error
	DeleteSubject(ctx context.Context, id uuid.UUID) error
}

type Event interface {
	ListEvents(ctx context.Context, filter *models.FilterEventDTO) ([]*models.Event, error)
	CreateEvent(ctx context.Context, event *models.Event) error
	UpdateEvent(ctx context.Context, event *models.Event) error
	DeleteEvent(ctx context.Context, eventID uuid.UUID) error
}

type Media interface {
	Upload(ctx context.Context, fileReader io.Reader, extension string) (url string, err error)
}

type Notifier interface {
	SendNotification(ctx context.Context, dto *models.SendNotificationDTO) error

	GetTemplates(ctx context.Context) ([]models.NotificationTemplate, error)
	CreateTemplate(ctx context.Context, template *models.NotificationTemplate) error
	UpdateTemplate(ctx context.Context, template *models.NotificationTemplate) error
	DeleteTemplate(ctx context.Context, templateID uuid.UUID) error
}

type Activity interface {
	ListActivitiesByStudent(ctx context.Context, studentID uuid.UUID) ([]models.Activity, error)
	CreateActivity(ctx context.Context, activity *models.Activity, studentID uuid.UUID) error
	UpdateActivity(ctx context.Context, activity *models.Activity, activityID uuid.UUID) error
	DeleteActivity(ctx context.Context, activityID uuid.UUID) error
}

type Assignment interface {
	ListAssignmentsBySubject(ctx context.Context, subjectID uuid.UUID) ([]models.Assignment, error)
	CreateAssignment(ctx context.Context, assignment *models.Assignment, subjectID uuid.UUID) error
	UpdateAssignment(ctx context.Context, assignment *models.Assignment) error
	DeleteAssignment(ctx context.Context, assignmentID uuid.UUID) error
	SubmitAssignment(ctx context.Context, studentID uuid.UUID, assignment *models.SubmittedAssigment) error
}

type Diet interface {
	GetDietByID(id uuid.UUID) (*models.Diet, error)
	ListDiets() ([]*models.Diet, error)
	CreateDiet(diet *models.Diet) error
	UpdateDiet(diet *models.Diet) error
	DeleteDiet(id uuid.UUID) error
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
	CreateDish(ctx context.Context, dish *models.Dish) error
	UpdateDish(ctx context.Context, dish *models.Dish) error
	DeleteDish(ctx context.Context, dishID uuid.UUID) error
	AddDietToDish(ctx context.Context, dishDiet *models.DishDiet) error
	AddTagToDish(ctx context.Context, dishTag *models.DishTag) error
}
