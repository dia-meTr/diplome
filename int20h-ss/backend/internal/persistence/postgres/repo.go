package postgres

import "oss-backend/internal/persistence"

func (p *Postgres) Auth() persistence.Auth {
	return p
}

func (p *Postgres) User() persistence.User {
	return p
}

func (p *Postgres) Faculty() persistence.Faculty {
	return p
}

func (p *Postgres) Group() persistence.Group {
	return p
}

func (p *Postgres) Subject() persistence.Subject {
	return p
}

func (p *Postgres) Event() persistence.Event {
	return p
}

func (p *Postgres) Notification() persistence.Notification {
	return p
}

func (p *Postgres) Activity() persistence.Activity {
	return p
}

func (p *Postgres) Assignment() persistence.Assignment {
	return p
}

func (p *Postgres) Product() persistence.Product {
	return p
}

func (p *Postgres) Diet() persistence.Diet {
	return p
}

func (p *Postgres) Tag() persistence.Tag {
	return p
}

func (p *Postgres) Dish() persistence.Dish {
	return p
}
