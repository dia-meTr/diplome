package models

import (
	"github.com/google/uuid"
)

type Dish struct {
	ID          uuid.UUID `json:"id" bun:",pk,type:uuid,default:uuid_generate_v4()"`
	Name        string    `json:"name"`
	PhotoUrl    string    `json:"photoUrl"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Tags        []*Tag    `json:"dishTags" bun:"m2m:dish_tags,join:Dish=Tag"`
	// DishProducts []*DishProduct `json:"dishProducts" bun:"m2m:dish_products,join:Dish=Product"`
}

type DishTag struct {
	DishID uuid.UUID `json:"dishId" bun:",pk"`
	Dish   *Dish     `bun:"rel:belongs-to,join:dish_id=id"`
	TagID  int       `json:"tagId" bun:",pk"`
	Tag    *Tag      `bun:"rel:belongs-to,join:tag_id=id"`
}

type DishProduct struct {
	DishID uuid.UUID `json:"dishId" bun:",pk"`
	TagID  uuid.UUID `json:"tagId" bun:",pk"`
	Amount int       `json:"amount"`
}
