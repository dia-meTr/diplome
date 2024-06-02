package models

import (
	"github.com/google/uuid"
)

type OrderStatus string

const (
	NewStatus        OrderStatus = "New"
	InProgressStatus OrderStatus = "In progress"
	ReadyStatus      OrderStatus = "Ready"
	DeliveredStatus  OrderStatus = "Delivered"
)

type Order struct {
	ID              uuid.UUID   `json:"id" bun:",pk,type:uuid,default:uuid_generate_v4()"`
	UserID          uuid.UUID   `json:"userId"`
	User            *User       `bun:"rel:belongs-to,join:user_id=id"`
	OrderDate       string      `json:"orderDate" bun:",nullzero"`
	DeliveryDate    string      `json:"deliveryDate" bun:",nullzero"`
	DeliveryAddress string      `json:"deliveryAddress"`
	OrderStatus     OrderStatus `json:"orderStatus"` // New, In progress, Ready, Delivered
}

type OrderDish struct {
	OrderID uuid.UUID `json:"orderId" bun:",pk"`
	Order   *Order    `bun:"rel:belongs-to,join:order_id=id"`
	DishID  uuid.UUID `json:"dishId" bun:",pk"`
	Dish    *Dish     `bun:"rel:belongs-to,join:dish_id=id"`
	Amount  int       `json:"amount"`
}
