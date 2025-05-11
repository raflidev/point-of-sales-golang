package service

import (
	"context"
	"golang-point-of-sales-system/exception"
	"golang-point-of-sales-system/helper"
	"golang-point-of-sales-system/modules/pembelian/domain/entity"
	"golang-point-of-sales-system/modules/pembelian/domain/repository"
	"golang-point-of-sales-system/modules/pembelian/dto/request"
	"golang-point-of-sales-system/modules/pembelian/dto/response"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type PembelianServiceImpl struct {
	PembelianRepository repository.PembelianRepository
	Validate            *validator.Validate
}

func NewPembelianService(pembelianRepository repository.PembelianRepository, validate *validator.Validate) PembelianService {
	return &PembelianServiceImpl{
		PembelianRepository: pembelianRepository,
		Validate:            validate,
	}
}

func (service *PembelianServiceImpl) Create(ctx context.Context, request request.PembelianCreateRequest) response.PembelianResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	pembelian := entity.Pembelian{
		Supplier_id: request.Supplier_id,
		Total_item:  request.Total_item,
		Total_harga: request.Total_harga,
		Diskon:      request.Diskon,
		Bayar:       request.Bayar,
	}

	pembelian = service.PembelianRepository.Save(ctx, pembelian)

	return helper.ToPembelianResponse(pembelian)
}

func (service *PembelianServiceImpl) Update(ctx context.Context, request request.PembelianUpdateRequest) response.PembelianResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	pembelian, err := service.PembelianRepository.FindById(ctx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	pembelian.Supplier_id = request.Supplier_id
	pembelian.Total_item = request.Total_item
	pembelian.Total_harga = request.Total_harga
	pembelian.Diskon = request.Diskon
	pembelian.Bayar = request.Bayar

	pembelian, err = service.PembelianRepository.Update(ctx, pembelian)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToPembelianResponse(pembelian)
}

func (service *PembelianServiceImpl) Delete(ctx context.Context, pembeliId uuid.UUID) {
	pembelian, err := service.PembelianRepository.FindById(ctx, pembeliId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.PembelianRepository.Delete(ctx, pembelian)
}

func (service *PembelianServiceImpl) FindById(ctx context.Context, pembeliId uuid.UUID) response.PembelianResponse {
	pembelian, err := service.PembelianRepository.FindById(ctx, pembeliId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToPembelianResponse(pembelian)
}

func (service *PembelianServiceImpl) FindAll(ctx context.Context) []response.PembelianResponse {
	pembelians := service.PembelianRepository.FindAll(ctx)
	var pembelianResponses []response.PembelianResponse
	for _, pembelian := range pembelians {
		pembelianResponses = append(pembelianResponses, helper.ToPembelianResponse(pembelian))
	}

	return pembelianResponses
}
