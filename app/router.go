package app

import (
	"golang-point-of-sales-system/adapters"
	"golang-point-of-sales-system/exception"
	categoryHandler "golang-point-of-sales-system/modules/categories/controller"
	memberHandler "golang-point-of-sales-system/modules/members/controller"
	pembelianHandler "golang-point-of-sales-system/modules/pembelian/controller"
	pembelianDetailHandler "golang-point-of-sales-system/modules/pembelianDetail/controller"
	productHandler "golang-point-of-sales-system/modules/products/controller"
	supplierHandler "golang-point-of-sales-system/modules/suppliers/controller"
	userHandler "golang-point-of-sales-system/modules/users/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	echo                      *echo.Echo
	productController         productHandler.ProductController
	supplierController        supplierHandler.SupplierController
	categoryController        categoryHandler.CategoryController
	userController            userHandler.UserController
	memberController          memberHandler.MemberController
	pembelianController       pembelianHandler.PembelianController
	pembelianDetailController pembelianDetailHandler.PembelianDetailController
}

func NewRouter(
	productController productHandler.ProductController,
	supplierController supplierHandler.SupplierController,
	categoryController categoryHandler.CategoryController,
	userController userHandler.UserController,
	memberController memberHandler.MemberController,
	pembelianController pembelianHandler.PembelianController,
	pembelianDetailController pembelianDetailHandler.PembelianDetailController,

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

	userGroup := apiV1.Group("/user")
	userGroup.GET("/lists", adapters.HttprouterHandlerToEchoHandler(userController.FindAll))
	userGroup.POST("/add", adapters.HttprouterHandlerToEchoHandler(userController.Create))
	userGroup.GET("/show/:userId", adapters.HttprouterHandlerToEchoHandler(userController.FindById))
	userGroup.PUT("/update/:userId", adapters.HttprouterHandlerToEchoHandler(userController.Update))
	userGroup.DELETE("/delete/:userId", adapters.HttprouterHandlerToEchoHandler(userController.Delete))
	userGroup.POST("/login", adapters.HttprouterHandlerToEchoHandler(userController.Login))
	userGroup.PUT("/change-password", adapters.HttprouterHandlerToEchoHandler(userController.ChangePassword))

	memberGroup := apiV1.Group("/member")
	memberGroup.GET("/lists", adapters.HttprouterHandlerToEchoHandler(memberController.FindAll))
	memberGroup.POST("/add", adapters.HttprouterHandlerToEchoHandler(memberController.Create))
	memberGroup.GET("/show/:memberId", adapters.HttprouterHandlerToEchoHandler(memberController.FindById))
	memberGroup.PUT("/update/:memberId", adapters.HttprouterHandlerToEchoHandler(memberController.Update))
	memberGroup.DELETE("/delete/:memberId", adapters.HttprouterHandlerToEchoHandler(memberController.Delete))

	pembelianGroup := apiV1.Group("/pembelian")
	pembelianGroup.GET("/lists", adapters.HttprouterHandlerToEchoHandler(pembelianController.FindAll))
	pembelianGroup.POST("/add", adapters.HttprouterHandlerToEchoHandler(pembelianController.Create))
	pembelianGroup.GET("/show/:pembelianId", adapters.HttprouterHandlerToEchoHandler(pembelianController.FindById))
	pembelianGroup.PUT("/update/:pembelianId", adapters.HttprouterHandlerToEchoHandler(pembelianController.Update))
	pembelianGroup.DELETE("/delete/:pembelianId", adapters.HttprouterHandlerToEchoHandler(pembelianController.Delete))

	pembelianDetailGroup := apiV1.Group("/pembelianDetail")
	pembelianDetailGroup.GET("/lists", adapters.HttprouterHandlerToEchoHandler(pembelianDetailController.FindAll))
	pembelianDetailGroup.POST("/add", adapters.HttprouterHandlerToEchoHandler(pembelianDetailController.Create))
	pembelianDetailGroup.GET("/show/:pembelianDetailId", adapters.HttprouterHandlerToEchoHandler(pembelianDetailController.FindById))
	pembelianDetailGroup.PUT("/update/:pembelianDetailId", adapters.HttprouterHandlerToEchoHandler(pembelianDetailController.Update))
	pembelianDetailGroup.DELETE("/delete/:pembelianDetailId", adapters.HttprouterHandlerToEchoHandler(pembelianDetailController.Delete))

	return router
}
