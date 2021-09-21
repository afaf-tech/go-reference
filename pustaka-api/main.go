package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "afaf",
			"bio":  "a software engineer and self-employed",
		})
	})
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "hello afaf",
			"bio":  "learing golang with agung setiawan mu",
		})
	})

	router.Run(":3000")
}
