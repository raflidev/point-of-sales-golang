package main

import (
	"fmt"
	"golang-point-of-sales-system/app"
	"golang-point-of-sales-system/controller"
	"golang-point-of-sales-system/helper"
	"golang-point-of-sales-system/middleware"
	"golang-point-of-sales-system/modules/products/domain/repository"
	"golang-point-of-sales-system/modules/products/domain/service"
	"net/http"

	_ "github.com/jmoiron/sqlx"

	"github.com/go-playground/validator/v10"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}

}

func main() {
	db := app.NewDB()
	validate := validator.New()
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository, validate)
	productController := controller.NewProductController(productService)
	router := app.NewRouter(productController)
	authMiddleware := middleware.NewAuthMiddleware(router)

	server := NewServer(authMiddleware)

	err := server.ListenAndServe()
	helper.PanicIfError(err)

	fmt.Println("Server is running on port 3000")
}
