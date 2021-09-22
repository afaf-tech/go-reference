package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "afaf",
		"bio":  "a software engineer and self-employed",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "hello afaf",
		"bio":  "learing golang with agung setiawan mu",
	})
}

func BookHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"name": "hello afaf " + id,
		"bio":  "learing golang with agung setiawan mu",
	})
}

func PostBookHandler(c *gin.Context) {
	var bookInput book.BookInput

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

func QueryHandler(c *gin.Context) {
	id := c.Query("id")
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{
		"name":  "hello afaf " + id,
		"price": price,
		"bio":   "learing golang with agung setiawan mu",
	})
}
