package main

import (
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
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	v1 := router.Group("/v1/")
	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/hello", bookHandler.HelloHandler)
	v1.GET("/books", bookHandler.GetAll)
	v1.GET("/books/:id", bookHandler.GetOne)
	v1.POST("/books", bookHandler.PostBookHandler)
	v1.GET("/query", bookHandler.QueryHandler)

	router.Run(":3000")
}
