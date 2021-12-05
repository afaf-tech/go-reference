//go:build wireinject
// +build wireinject

// go:build wireinject
package main

import (
	"afaf-tech/belajar-golang-restful-api/app"
	"afaf-tech/belajar-golang-restful-api/controller"
	"afaf-tech/belajar-golang-restful-api/middleware"
	"afaf-tech/belajar-golang-restful-api/repository"
	"afaf-tech/belajar-golang-restful-api/service"
	"net/http"

	validator "github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializeServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
