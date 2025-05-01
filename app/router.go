package app

import (
	"golang-point-of-sales-system/exception"
	productHandler "golang-point-of-sales-system/modules/products/controller"
	supplierHandler "golang-point-of-sales-system/modules/suppliers/controller"

	"github.com/julienschmidt/httprouter"
)

type Router struct {
	router             *httprouter.Router
	productController  productHandler.ProductController
	supplierController supplierHandler.SupplierController
}

func NewRouter(
	productController productHandler.ProductController,
	supplierController supplierHandler.SupplierController,
) *httprouter.Router {
	router := httprouter.New()

	// Product routes
	router.GET("/api/v1/product/lists", productController.FindAll)
	router.POST("/api/v1/product/add", productController.Create)
	router.GET("/api/v1/product/show/:productId", productController.FindById)
	router.PUT("/api/v1/product/update/:productId", productController.Update)
	router.DELETE("/api/v1/product/delete/:productId", productController.Delete)

	// Supplier routes
	router.GET("/api/v1/supplier/lists", supplierController.FindAll)
	router.POST("/api/v1/supplier/add", supplierController.Create)
	router.GET("/api/v1/supplier/show/:supplierId", supplierController.FindById)
	router.PUT("/api/v1/supplier/update/:supplierId", supplierController.Update)
	router.DELETE("/api/v1/supplier/delete/:supplierId", supplierController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
