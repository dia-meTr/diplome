package postgres

import "oss-backend/internal/persistence"

func (p *Postgres) Auth() persistence.Auth {
	return p
}

func (p *Postgres) User() persistence.User {
	return p
}

func (p *Postgres) Product() persistence.Product {
	return p
}

func (p *Postgres) Tag() persistence.Tag {
	return p
}

func (p *Postgres) Dish() persistence.Dish {
	return p
}

func (p *Postgres) Order() persistence.Order {
	return p
}

func (p *Postgres) ShoppingCart() persistence.ShoppingCard {
	return p
}
