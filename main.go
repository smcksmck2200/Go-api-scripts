package main

import (
	"GOLANG_WEBAPP/book"
	"GOLANG_WEBAPP/handler"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=password1 dbname=postgres port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("db connection error")
	}
	fmt.Println("Successfully connected!")

	db.AutoMigrate(&book.Book{})
	////CRUD
	//book := book.Book{}
	//book.Title = "The Science Of Getting Rich"
	//book.Price = 50
	//book.Discount = 10
	//book.Rating = 5
	//book.Description = "A book of self mastery"
	//
	//err = db.Create(&book).Error
	//if err != nil {
	//	fmt.Println("==========================")
	//	fmt.Println("Error creating book record")
	//	fmt.Println("==========================")
	//}
	var book book.Book

	err = db.Debug().Where("id = ?", 1).First(&book).Error
	if err != nil {
		fmt.Println("==========================")
		fmt.Println("Error finding book record")
		fmt.Println("==========================")
	}
	err = db.Delete(&book).Error
	if err != nil {
		fmt.Println("==========================")
		fmt.Println("Error deleting book record")
		fmt.Println("==========================")
	}
	// ============
	// update data
	// ============
	//book.Title = "Richest Man In Babylon(revised edition)"
	//err = db.Save(&book).Error
	//if err != nil {
	//	fmt.Println("==========================")
	//	fmt.Println("Error updating book record")
	//	fmt.Println("==========================")
	//}
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run(":3000")
}
