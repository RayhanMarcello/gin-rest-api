package main

import (
	"belajar-golang/handler"
	"belajar-golang/product"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()

	dsn := "app:root@tcp(127.0.0.1:3306)/belajar_golang?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		fmt.Println("connection db err")
	}

	fmt.Println("db connection sucsess")

	db.AutoMigrate(&product.Product{})

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/products", handler.ProdHandler)
	v1.GET("/products/:id", handler.ProdIdHandler)
	v1.GET("/querry",handler.QuerryHandler)
	v1.POST("/products", handler.PostProductHandler)


	router.Run(":9000")
} 





