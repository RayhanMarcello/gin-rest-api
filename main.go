package main

import (
	"belajar-golang/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/products", handler.ProdHandler)
	v1.GET("/products/:id", handler.ProdIdHandler)
	v1.GET("/querry",handler.QuerryHandler)
	v1.POST("/products", handler.PostProductHandler)


	router.Run(":9000")
} 





