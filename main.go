package main

import (
	"fmt"
	"golang-point-of-sales-system/app"
	"golang-point-of-sales-system/controller"
	"golang-point-of-sales-system/helper"
	"golang-point-of-sales-system/middleware"
	productRepo "golang-point-of-sales-system/modules/products/domain/repository"
	productService "golang-point-of-sales-system/modules/products/domain/service"
	supplierRepo "golang-point-of-sales-system/modules/suppliers/domain/repository"
	supplierService "golang-point-of-sales-system/modules/suppliers/domain/service"
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
	productRepository := productRepo.NewProductRepository(db)
	productService := productService.NewProductService(productRepository, validate)
	productController := controller.NewProductController(productService)

	supplierRepository := supplierRepo.NewSupplierRepository(db)
	supplierService := supplierService.NewSupplierService(supplierRepository, validate)
	supplierController := controller.NewSupplierController(supplierService)
	router := app.NewRouter(productController, supplierController)
	authMiddleware := middleware.NewAuthMiddleware(router)

	server := NewServer(authMiddleware)

	err := server.ListenAndServe()
	helper.PanicIfError(err)

	fmt.Println("Server is running on port 3000")
}
