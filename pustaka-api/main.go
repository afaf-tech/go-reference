package main

import (
	"fmt"
	"log"
	"pustaka-api/app"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	db := app.GetConnection()
	// db.AutoMigrate(&book.Book{})

	router := gin.Default()

	bookRepository := book.NewRepository(db)

	book := book.Book{
		Title:       "The $100 STARTUP",
		Description: "The $100 STARTUP",
		Rating:      4,
		Discount:    0,
	}
	_, err := bookRepository.Create(book)
	if err != nil {
		log.Fatal(err)
	}

	books, err := bookRepository.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(books)

	v1 := router.Group("/v1/")
	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id", handler.BookHandler)
	v1.POST("/books", handler.PostBookHandler)
	v1.GET("/query", handler.QueryHandler)

	router.Run(":3000")
}
