package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"afaf-tech/belajar-golang-restful-api/app"
	"afaf-tech/belajar-golang-restful-api/controller"
	"afaf-tech/belajar-golang-restful-api/helper"
	"afaf-tech/belajar-golang-restful-api/middleware"
	"afaf-tech/belajar-golang-restful-api/repository"
	"afaf-tech/belajar-golang-restful-api/service"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	fmt.Println("Server listens on port 3000")

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
