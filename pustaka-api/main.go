package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
}

func postBookHandler(c *gin.Context) {
	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s ", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
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
