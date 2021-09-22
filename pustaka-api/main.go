package main

import (
	"fmt"
	"pustaka-api/app"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	db := app.GetConnection()
	db.AutoMigrate(&book.Book{})

	router := gin.Default()

	// Create Data
	// book := &book.Book{
	// 	Title:       "$100 STARTUP",
	// 	Description: "Author : Chris Guillebeau",
	// 	Rating:      4,
	// 	Discount:    2,
	// }

	// err := db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("===================")
	// 	fmt.Println("Error Creating Book")
	// 	fmt.Println("===================")
	// }

	// get one data
	// var book book.Book
	// err := db.First(&book, 2).Error
	// if err != nil {
	// 	fmt.Println("===================")
	// 	fmt.Println("Error Query Book")
	// 	fmt.Println("===================")
	// }
	// fmt.Println(book)

	//  GET ALL DATA
	var books []book.Book
	err := db.Where("Rating = ?", "4").Find(&books).Error
	if err != nil {
		fmt.Println("===================")
		fmt.Println("Error Query Book")
		fmt.Println("===================")
	}
	fmt.Println(books)

	// var books []book.Book
	// err := db.Find(&books).Error
	// if err != nil {
	// 	fmt.Println("===================")
	// 	fmt.Println("Error Query Book")
	// 	fmt.Println("===================")
	// }
	// for _, book := range books {
	// 	fmt.Println(book)
	// }

	// UPDATE DATA
	// var book book.Book

	// err = db.Debug().Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("===================")
	// 	fmt.Println("Error Query Book")
	// 	fmt.Println("===================")
	// }
	// book.Title = "Niana {Revised Edition}"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("===================")
	// 	fmt.Println("Error Query Book")
	// 	fmt.Println("===================")
	// }

	// DELETE DATA
	var book book.Book

	err = db.Debug().Where("id = ?", 1).First(&book).Error
	if err != nil {
		fmt.Println("===================")
		fmt.Println("Error Query Book")
		fmt.Println("===================")
	}
	err = db.Delete(&book).Error
	if err != nil {
		fmt.Println("===================")
		fmt.Println("Error Query Book")
		fmt.Println("===================")
	}

	v1 := router.Group("/v1/")
	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id", handler.BookHandler)
	v1.POST("/books", handler.PostBookHandler)
	v1.GET("/query", handler.QueryHandler)

	router.Run(":3000")
}
