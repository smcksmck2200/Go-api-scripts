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

	bookRepository := book.NewRepository(db)

	bookService := book.NewService(bookRepository)

	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/books", bookHandler.GetBooks)
	v1.GET("/books/:id", bookHandler.GetBook)
	v1.POST("/books", bookHandler.CreateBook)
	v1.PUT("/books/:id", bookHandler.UpdateBook)
	v1.DELETE("/books/:id", bookHandler.DeleteBook)
	router.Run()
}

//main
//handler
//service
//repository
//db
//postgresSQL
