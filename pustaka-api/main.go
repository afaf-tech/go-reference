package main

import (
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1/")
	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id", handler.BookHandler)
	v1.POST("/books", handler.PostBookHandler)
	v1.GET("/query", handler.QueryHandler)

	router.Run(":3000")
}
