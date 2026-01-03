package main

import (
	// "belajar-golang/handler"
	// "belajar-golang/product"
	// "fmt"
	"belajar-golang/database"
	"belajar-golang/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.ConnectionDatabase()

	//CRUD
	// prod := product.Product{}
	// prod.Name = "cireng2"
	// prod.Price = 200002
	// prod.Description = "cireng makanan yang enak 2"
	// prod.Rating = 52
	// err = db.Create(&prod).Error
	// if err != nil{
	// 	fmt.Println("cannot create to db")
	// }

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/products", handler.ProdHandler)
	v1.GET("/products/:id", handler.ProdIdHandler)
	v1.GET("/querry",handler.QuerryHandler)
	v1.POST("/products", handler.PostProductHandler)


	router.Run(":9000")
} 





