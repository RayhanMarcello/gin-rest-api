package database

import (
	"belajar-golang/product"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDatabase() {
	dsn := "app:root@tcp(127.0.0.1:3306)/belajar_golang?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		fmt.Println("connection db err")
	}
	fmt.Println("db connection sucsess")

	db.AutoMigrate(&product.Product{})

	DB = db
}