package main

import (
	"afaf-tech/belajar-golang-restful-api/helper"
	"afaf-tech/belajar-golang-restful-api/middleware"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}
func main() {

	server := InitializeServer()

	fmt.Println("Server listens on port 3000")

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
