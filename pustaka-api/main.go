package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/hello", helloHandler)
	router.GET("/books/:id", bookHandler)
	router.POST("/books", postBookHandler)
	router.GET("/query", queryHandler)

	router.Run(":3000")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "afaf",
		"bio":  "a software engineer and self-employed",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "hello afaf",
		"bio":  "learing golang with agung setiawan mu",
	})
}

func bookHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"name": "hello afaf " + id,
		"bio":  "learing golang with agung setiawan mu",
	})
}

type BookInput struct {
	Title string `json:"title"`
	Price int    `json:"price"`
}

func postBookHandler(c *gin.Context) {
	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, bookInput)
}

func queryHandler(c *gin.Context) {
	id := c.Query("id")
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{
		"name":  "hello afaf " + id,
		"price": price,
		"bio":   "learing golang with agung setiawan mu",
	})
}
