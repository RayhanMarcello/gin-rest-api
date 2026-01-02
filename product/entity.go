package product

import "time"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       int
	Rating      int
	CreateAt    time.Time
	UpdateAt	time.Time
}