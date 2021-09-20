package main

import (
	"belajar-golang-restful-api/app"
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/repository"
	"belajar-golang-restful-api/service"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func main() {
	validate := validator.New()
	db := app.NewDB()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := httprouter.New()

	router.GET("/api/cateogories", categoryController.FindAll)
	router.GET("/api/cateogories/:categoryId", categoryController.FindById)
	router.POST("/api/cateogories", categoryController.Create)
	router.PUT("/api/cateogories/:categoryId", categoryController.Update)
	router.DELETE("/api/cateogories/:categoryId", categoryController.Delete)
}
