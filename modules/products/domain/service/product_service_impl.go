package service

import (
	"context"
	"golang-point-of-sales-system/helper"
	"golang-point-of-sales-system/modules/products/domain/entity"
	"golang-point-of-sales-system/modules/products/domain/repository"
	"golang-point-of-sales-system/modules/products/dto/request"
	"golang-point-of-sales-system/modules/products/dto/response"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sqlx.DB
	Validate          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, db *sqlx.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                db,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request request.ProductCreateRequest) response.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product := entity.Product{
		Kode_produk: request.Kode_produk,
		Nama_produk: request.Nama_produk,
		Merk:        request.Merk,
		Harga_beli:  request.Harga_beli,
		Harga_jual:  request.Harga_jual,
		Stok:        request.Stok,
	}

	product = service.ProductRepository.Save(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) Update(ctx context.Context, request request.ProductUpdateRequest) response.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	product.Kode_produk = request.Kode_produk
	product.Nama_produk = request.Nama_produk
	product.Merk = request.Merk
	product.Harga_beli = request.Harga_beli
	product.Harga_jual = request.Harga_jual
	product.Stok = request.Stok

	product = service.ProductRepository.Update(ctx, tx, product)

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) Delete(ctx context.Context, productId int) {
	tx, err := service.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	helper.PanicIfError(err)

	service.ProductRepository.Delete(ctx, tx, product)
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId int) response.ProductResponse {
	tx, err := service.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	helper.PanicIfError(err)

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []response.ProductResponse {
	tx, err := service.DB.Beginx()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepository.FindAll(ctx, tx)
	return helper.ToProductResponses(products)
}
