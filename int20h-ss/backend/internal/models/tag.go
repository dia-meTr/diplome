package models

import "github.com/google/uuid"

type Tag struct {
	ID   uuid.UUID `json:"id" bun:",pk,nullzero"`
	Name string    `json:"name"`
}
