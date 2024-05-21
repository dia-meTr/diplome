package models

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID `json:"id" bun:",pk,nullzero"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
