package app

import (
	"golang-point-of-sales-system/controller"
	"golang-point-of-sales-system/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(productController controller.ProductController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/v1/product/lists", productController.FindAll)
	router.POST("/api/v1/product/add", productController.Create)
	router.GET("/api/v1/product/show/:productId", productController.FindById)
	router.PUT("/api/v1/product/update/:productId", productController.Update)
	router.DELETE("/api/v1/product/delete/:productId", productController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
