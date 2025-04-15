package service

import (
	"context"
	"golang-point-of-sales-system/exception"
	"golang-point-of-sales-system/helper"
	"golang-point-of-sales-system/modules/products/domain/entity"
	"golang-point-of-sales-system/modules/products/domain/repository"
	"golang-point-of-sales-system/modules/products/dto/request"
	"golang-point-of-sales-system/modules/products/dto/response"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	Validate          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request request.ProductCreateRequest) response.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	product := entity.Product{
		Kode_produk: request.Kode_produk,
		Nama_produk: request.Nama_produk,
		Merk:        request.Merk,
		Harga_beli:  request.Harga_beli,
		Harga_jual:  request.Harga_jual,
		Stok:        request.Stok,
	}

	product = service.ProductRepository.Save(ctx, product)

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) Update(ctx context.Context, request request.ProductUpdateRequest) response.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	product, err := service.ProductRepository.FindById(ctx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	product.Kode_produk = request.Kode_produk
	product.Nama_produk = request.Nama_produk
	product.Merk = request.Merk
	product.Harga_beli = request.Harga_beli
	product.Harga_jual = request.Harga_jual
	product.Stok = request.Stok

	updatedProduct, err := service.ProductRepository.Update(ctx, product)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToProductResponse(updatedProduct)
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId uuid.UUID) {

	product, err := service.ProductRepository.FindById(ctx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.ProductRepository.Delete(ctx, product)
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId uuid.UUID) response.ProductResponse {

	product, err := service.ProductRepository.FindById(ctx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []response.ProductResponse {
	products := service.ProductRepository.FindAll(ctx)
	return helper.ToProductResponses(products)
}
