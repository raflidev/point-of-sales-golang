package main

import (
	"fmt"
	"golang-point-of-sales-system/app"
	"golang-point-of-sales-system/helper"
	"golang-point-of-sales-system/middleware"
	categoryController "golang-point-of-sales-system/modules/categories/controller"
	categoryRepo "golang-point-of-sales-system/modules/categories/domain/repository"
	categoryService "golang-point-of-sales-system/modules/categories/domain/service"
	memberController "golang-point-of-sales-system/modules/members/controller"
	memberRepo "golang-point-of-sales-system/modules/members/domain/repository"
	memberService "golang-point-of-sales-system/modules/members/domain/service"
	pembelianController "golang-point-of-sales-system/modules/pembelian/controller"
	pembelianRepo "golang-point-of-sales-system/modules/pembelian/domain/repository"
	pembelianService "golang-point-of-sales-system/modules/pembelian/domain/service"
	pembelianDetailController "golang-point-of-sales-system/modules/pembelianDetail/controller"
	pembelianDetailRepo "golang-point-of-sales-system/modules/pembelianDetail/domain/repository"
	pembelianDetailService "golang-point-of-sales-system/modules/pembelianDetail/domain/service"
	penjualanController "golang-point-of-sales-system/modules/penjualan/controller"
	penjualanRepo "golang-point-of-sales-system/modules/penjualan/domain/repository"
	penjualanService "golang-point-of-sales-system/modules/penjualan/domain/service"
	productController "golang-point-of-sales-system/modules/products/controller"
	productRepo "golang-point-of-sales-system/modules/products/domain/repository"
	productService "golang-point-of-sales-system/modules/products/domain/service"
	supplierController "golang-point-of-sales-system/modules/suppliers/controller"
	supplierRepo "golang-point-of-sales-system/modules/suppliers/domain/repository"
	supplierService "golang-point-of-sales-system/modules/suppliers/domain/service"
	userController "golang-point-of-sales-system/modules/users/controller"
	userRepo "golang-point-of-sales-system/modules/users/domain/repository"
	userService "golang-point-of-sales-system/modules/users/domain/service"

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
	productHandler := productController.NewProductController(productService)

	supplierRepository := supplierRepo.NewSupplierRepository(db)
	supplierService := supplierService.NewSupplierService(supplierRepository, validate)
	supplierHandler := supplierController.NewSupplierController(supplierService)

	categoryRepository := categoryRepo.NewCategoryRepository(db)
	categoryService := categoryService.NewCategoryService(categoryRepository, validate)
	categoryHandler := categoryController.NewCategoryController(categoryService)

	userRepository := userRepo.NewUserRepository(db)
	userService := userService.NewUserService(userRepository, validate)
	userHandler := userController.NewUserController(userService)

	memberRepository := memberRepo.NewMemberRepository(db)
	memberService := memberService.NewMemberService(memberRepository, validate)
	memberHandler := memberController.NewMemberController(memberService)

	pembelianRepository := pembelianRepo.NewPembelianRepository(db)
	pembelianService := pembelianService.NewPembelianService(pembelianRepository, validate)
	pembelianHandler := pembelianController.NewPembelianController(pembelianService)

	pembelianDetailRepository := pembelianDetailRepo.NewPembelianDetailRepository(db)
	pembelianDetailService := pembelianDetailService.NewPembelianDetailService(pembelianDetailRepository, validate)
	pembelianDetailHandler := pembelianDetailController.NewPembelianDetailController(pembelianDetailService)

	penjualanRepository := penjualanRepo.NewPenjualanRepository(db)
	penjualanService := penjualanService.NewPenjualanService(penjualanRepository, validate)
	penjualanHandler := penjualanController.NewPenjualanController(penjualanService)

	router := app.NewRouter(productHandler, supplierHandler, categoryHandler, userHandler, memberHandler, pembelianHandler, pembelianDetailHandler, penjualanHandler)
	authMiddleware := middleware.NewAuthMiddleware(router)

	server := NewServer(authMiddleware)

	err := server.ListenAndServe()
	helper.PanicIfError(err)

	fmt.Println("Server is running on port 3000")
}
