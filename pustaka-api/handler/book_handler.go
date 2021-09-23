package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService: bookService}
}

func (handler *bookHandler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "afaf",
		"bio":  "a software engineer and self-employed",
	})
}

func (handler *bookHandler) HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "hello afaf",
		"bio":  "learing golang with agung setiawan mu",
	})
}

func (handler *bookHandler) GetOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("error conver id not number")
		c.JSON(http.StatusBadRequest, "error conver id not number")
		return
	}
	book, _ := handler.bookService.FindOne(id)
	c.JSON(http.StatusOK, book)
}

func (handler *bookHandler) GetAll(c *gin.Context) {
	books, _ := handler.bookService.FindAll()

	c.JSON(http.StatusOK, books)
}

func (handler *bookHandler) PostBookHandler(c *gin.Context) {
	var bookInput book.BookRequest

	err := c.ShouldBindJSON(&bookInput)

	if err != nil {
		serr, ok := err.(validator.ValidationErrors)
		if !ok {
			errorMessages := []string{}
			for _, e := range serr {
				errorMessage := fmt.Sprintf("Error on field %s, condition %s ", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
		}
		return
	}

	book, err := handler.bookService.Create(bookInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	c.JSON(http.StatusOK, book)
}

func (handler *bookHandler) QueryHandler(c *gin.Context) {
	id := c.Query("id")
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{
		"name":  "hello afaf " + id,
		"price": price,
		"bio":   "learing golang with agung setiawan mu",
	})
}
