package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          int 	
	Name        string	`json:"name"`
	Description string	`json:"description"`
	Price       int		`json:"price"`
	Rating      int		`json:"rating"`
}