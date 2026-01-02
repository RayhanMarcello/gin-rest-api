package handler

import (
	"net/http"

	"belajar-golang/product"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "rayhan marcello",
	})
}

func ProdHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"product": "cireng",
	})
}

func ProdIdHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id-product": id,
	})
}

func QuerryHandler(c *gin.Context) {
	title := c.Query("title")
	c.JSON(http.StatusOK, gin.H{
		"title": title,
	})
}

func PostProductHandler(c *gin.Context) {
	var productInput product.ProductInput
	err := c.ShouldBindBodyWithJSON(&productInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err-message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":  productInput.Name,
		"price": productInput.Price,
	})
}

