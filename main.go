package main

import (
	"golang-point-of-sales-system/app"
	"golang-point-of-sales-system/controller"
	"golang-point-of-sales-system/exception"
	"golang-point-of-sales-system/helper"
	"golang-point-of-sales-system/modules/products/domain/repository"
	"golang-point-of-sales-system/modules/products/domain/service"
	"net/http"

	_ "github.com/jmoiron/sqlx"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository, validate)
	productController := controller.NewProductController(productService)

	router := httprouter.New()

	router.GET("/api/v1/product/lists", productController.FindAll)
	router.POST("/api/v1/product/add", productController.Create)
	router.GET("/api/v1/product/show/:productId", productController.FindById)
	router.PUT("/api/v1/product/update/:productId", productController.Update)
	router.DELETE("/api/v1/product/delete/:productId", productController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
