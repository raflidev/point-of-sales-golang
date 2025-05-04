package app

import (
	"golang-point-of-sales-system/adapters"
	"golang-point-of-sales-system/exception"
	categoryHandler "golang-point-of-sales-system/modules/categories/controller"
	productHandler "golang-point-of-sales-system/modules/products/controller"
	supplierHandler "golang-point-of-sales-system/modules/suppliers/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	echo               *echo.Echo
	productController  productHandler.ProductController
	supplierController supplierHandler.SupplierController
}

func NewRouter(
	productController productHandler.ProductController,
	supplierController supplierHandler.SupplierController,
	categoryController categoryHandler.CategoryController,
) *echo.Echo {
	router := echo.New()

	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	// Adaptasi error handler dari httprouter ke echo
	router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					exception.ErrorHandler(c.Response().Writer, c.Request(), r)
				}
			}()
			return next(c)
		}
	})

	// Group utama untuk versi API
	apiV1 := router.Group("/api/v1")

	// Group untuk produk
	productGroup := apiV1.Group("/product")
	productGroup.GET("/lists", adapters.HttprouterHandlerToEchoHandler(productController.FindAll))
	productGroup.POST("/add", adapters.HttprouterHandlerToEchoHandler(productController.Create))
	productGroup.GET("/show/:productId", adapters.HttprouterHandlerToEchoHandler(productController.FindById))
	productGroup.PUT("/update/:productId", adapters.HttprouterHandlerToEchoHandler(productController.Update))
	productGroup.DELETE("/delete/:productId", adapters.HttprouterHandlerToEchoHandler(productController.Delete))

	// Group untuk supplier
	supplierGroup := apiV1.Group("/supplier")
	supplierGroup.GET("/lists", adapters.HttprouterHandlerToEchoHandler(supplierController.FindAll))
	supplierGroup.POST("/add", adapters.HttprouterHandlerToEchoHandler(supplierController.Create))
	supplierGroup.GET("/show/:supplierId", adapters.HttprouterHandlerToEchoHandler(supplierController.FindById))
	supplierGroup.PUT("/update/:supplierId", adapters.HttprouterHandlerToEchoHandler(supplierController.Update))
	supplierGroup.DELETE("/delete/:supplierId", adapters.HttprouterHandlerToEchoHandler(supplierController.Delete))

	categoryGroup := apiV1.Group("/category")
	categoryGroup.GET("/lists", adapters.HttprouterHandlerToEchoHandler(categoryController.FindAll))
	categoryGroup.POST("/add", adapters.HttprouterHandlerToEchoHandler(categoryController.Create))
	categoryGroup.GET("/show/:categoryId", adapters.HttprouterHandlerToEchoHandler(categoryController.FindById))
	categoryGroup.PUT("/update/:categoryId", adapters.HttprouterHandlerToEchoHandler(categoryController.Update))
	categoryGroup.DELETE("/delete/:categoryId", adapters.HttprouterHandlerToEchoHandler(categoryController.Delete))

	return router
}
