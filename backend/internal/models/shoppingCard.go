package models

import (
	"github.com/google/uuid"
)

type ShoppingCard struct {
	UserID uuid.UUID `json:"userId" bun:",pk"`
	User   *User     `bun:"rel:belongs-to,join:user_id=id"`
	DishID uuid.UUID `json:"dishId" bun:",pk"`
	Dish   *Dish     `bun:"rel:belongs-to,join:dish_id=id"`
	Amount int       `json:"amount"`
}
