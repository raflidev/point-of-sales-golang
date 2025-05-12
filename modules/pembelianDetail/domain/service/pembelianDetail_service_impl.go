package service

import (
	"context"
	"golang-point-of-sales-system/exception"
	"golang-point-of-sales-system/helper"
	"golang-point-of-sales-system/modules/pembelianDetail/domain/entity"
	"golang-point-of-sales-system/modules/pembelianDetail/domain/repository"
	"golang-point-of-sales-system/modules/pembelianDetail/dto/request"
	"golang-point-of-sales-system/modules/pembelianDetail/dto/response"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type PembelianDetailServiceImpl struct {
	PembelianDetailRepository repository.PembelianDetailRepository
	Validate                  *validator.Validate
}

func NewPembelianDetailService(pembelianDetailRepository repository.PembelianDetailRepository, validate *validator.Validate) PembelianDetailService {
	return &PembelianDetailServiceImpl{
		PembelianDetailRepository: pembelianDetailRepository,
		Validate:                  validate,
	}
}

func (service *PembelianDetailServiceImpl) Create(ctx context.Context, request request.PembelianDetailCreateRequest) response.PembelianDetailResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	pembelianDetail := entity.PembelianDetail{
		Pembelian_id: request.Pembelian_id,
		Produk_id:    request.Produk_id,
		Harga_beli:   request.Harga_beli,
		Jumlah:       request.Jumlah,
		Subtotal:     request.Subtotal,
	}

	pembelianDetail = service.PembelianDetailRepository.Save(ctx, pembelianDetail)

	return helper.ToPembelianDetailResponse(pembelianDetail)
}

func (service *PembelianDetailServiceImpl) Update(ctx context.Context, request request.PembelianDetailUpdateRequest) response.PembelianDetailResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	pembelianDetail, err := service.PembelianDetailRepository.FindById(ctx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	pembelianDetail.Pembelian_id = request.Pembelian_id
	pembelianDetail.Produk_id = request.Produk_id
	pembelianDetail.Harga_beli = request.Harga_beli
	pembelianDetail.Jumlah = request.Jumlah
	pembelianDetail.Subtotal = request.Subtotal

	pembelianDetail, err = service.PembelianDetailRepository.Update(ctx, pembelianDetail)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToPembelianDetailResponse(pembelianDetail)
}

func (service *PembelianDetailServiceImpl) Delete(ctx context.Context, pembelianDetailId uuid.UUID) {
	pembelianDetail, err := service.PembelianDetailRepository.FindById(ctx, pembelianDetailId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.PembelianDetailRepository.Delete(ctx, pembelianDetail)
}

func (service *PembelianDetailServiceImpl) FindById(ctx context.Context, pembelianDetailId uuid.UUID) response.PembelianDetailResponse {
	pembelianDetail, err := service.PembelianDetailRepository.FindById(ctx, pembelianDetailId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToPembelianDetailResponse(pembelianDetail)
}

func (service *PembelianDetailServiceImpl) FindAll(ctx context.Context) []response.PembelianDetailResponse {
	datas := service.PembelianDetailRepository.FindAll(ctx)
	var pembelianDetailResponses []response.PembelianDetailResponse
	for _, pembelianDetail := range datas {
		pembelianDetailResponses = append(pembelianDetailResponses, helper.ToPembelianDetailResponse(pembelianDetail))
	}

	return pembelianDetailResponses
}
