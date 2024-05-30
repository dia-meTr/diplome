package models

import "github.com/google/uuid"

type Role string

const (
	RoleAdmin   Role = "customer"
	RoleTeacher Role = "admin"
	RoleStudent Role = "delivery_man"
)

type User struct {
	ID        uuid.UUID `json:"id" bun:",pk,nullzero"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	UserRole  Role      `json:"role" bun:",nullzero"`
}
