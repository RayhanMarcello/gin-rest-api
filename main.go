package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/products", prodHandler)
	router.GET("/products/:id", prodIdHandler)
	router.GET("/querry",querryHandler)

	router.POST("/products", postProductHandler)

	router.Run(":9000")
} 

func rootHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
			"name": "rayhan marcello",
		})
}

func prodHandler(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
			"product" : "cireng", 
		})
}

func prodIdHandler(c *gin.Context){
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id-product" : id,
	})
}

func querryHandler(c *gin.Context){
	title := c.Query("title")
	c.JSON(http.StatusOK, gin.H{
		"title" : title,
	})
}

type ProductInput struct{
	Name string	`json:"name" binding:"required"`
	Price int	`json:"price" binding:"required,number"` 
}

func postProductHandler(c *gin.Context){
	var productInput ProductInput
	err := c.ShouldBindBodyWithJSON(&productInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err-message" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name" : productInput.Name,
		"price" : productInput.Price,
	})
}


