package product

type ProductInput struct {
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price" binding:"required,number"`
}
