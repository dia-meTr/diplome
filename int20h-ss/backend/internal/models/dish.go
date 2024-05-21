package models

import (
	"github.com/google/uuid"
)

type Dish struct {
	ID          uuid.UUID `json:"id" bun:",pk,type:uuid,default:uuid_generate_v4()"`
	Name        string    `json:"name"`
	PhotoUrl    string    `json:"photoUrl"`
	Description string    `json:"description"`
}

type DishTag struct {
	DishID uuid.UUID `json:"dishId"`
	TagID  uuid.UUID `json:"tagId"`
}

type DishDiet struct {
	DishID uuid.UUID `json:"dishId"`
	DietID uuid.UUID `json:"dietId"`
}

type DishProduct struct {
	DishID uuid.UUID `json:"dishId"`
	TagID  uuid.UUID `json:"tagId"`
	Amount float64   `json:"amount"`
}
