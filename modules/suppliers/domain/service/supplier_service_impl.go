package service

import (
	"context"
	"golang-point-of-sales-system/exception"
	"golang-point-of-sales-system/helper"
	"golang-point-of-sales-system/modules/suppliers/domain/entity"
	"golang-point-of-sales-system/modules/suppliers/domain/repository"
	"golang-point-of-sales-system/modules/suppliers/dto/request"
	"golang-point-of-sales-system/modules/suppliers/dto/response"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type SupplierServiceImpl struct {
	SupplierRepository repository.SupplierRepository
	Validate           *validator.Validate
}

func NewSupplierService(supplierRepository repository.SupplierRepository, validate *validator.Validate) SupplierService {
	return &SupplierServiceImpl{
		SupplierRepository: supplierRepository,
		Validate:           validate,
	}
}

func (service *SupplierServiceImpl) Create(ctx context.Context, request request.SupplierCreateRequest) response.SupplierResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	supplier := entity.Supplier{
		Nama:    request.Nama,
		Alamat:  request.Alamat,
		Telepon: request.Telepon,
	}

	supplier = service.SupplierRepository.Save(ctx, supplier)

	return helper.ToSupplierResponse(supplier)
}

func (service *SupplierServiceImpl) Update(ctx context.Context, request request.SupplierUpdateRequest) response.SupplierResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	supplier, err := service.SupplierRepository.FindById(ctx, request.Id)
	if err != nil {
		panic(err)
	}

	supplier.Nama = request.Nama
	supplier.Alamat = request.Alamat
	supplier.Telepon = request.Telepon

	supplier, err = service.SupplierRepository.Update(ctx, supplier)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToSupplierResponse(supplier)
}

func (service *SupplierServiceImpl) Delete(ctx context.Context, id uuid.UUID) {
	supplier, err := service.SupplierRepository.FindById(ctx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.SupplierRepository.Delete(ctx, supplier)
}

func (service *SupplierServiceImpl) FindById(ctx context.Context, id uuid.UUID) response.SupplierResponse {
	supplier, err := service.SupplierRepository.FindById(ctx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToSupplierResponse(supplier)
}

func (service *SupplierServiceImpl) FindAll(ctx context.Context) []response.SupplierResponse {
	suppliers := service.SupplierRepository.FindAll(ctx)
	if len(suppliers) == 0 {
		panic(exception.NewNotFoundError("Supplier not found"))
	}

	return helper.ToSupplierResponses(suppliers)
}
