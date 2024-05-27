package models

type Tag struct {
	ID     int    `json:"id" bun:",pk,nullzero"`
	Name   string `json:"name"`
	Dishes []Dish `bun:"m2m:dish_tags,join:Tag=Dish"`
}
