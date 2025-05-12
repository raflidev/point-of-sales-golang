package service

import (
	"context"
	"golang-point-of-sales-system/exception"
	"golang-point-of-sales-system/helper"
	"golang-point-of-sales-system/modules/penjualan/domain/entity"
	"golang-point-of-sales-system/modules/penjualan/domain/repository"
	"golang-point-of-sales-system/modules/penjualan/dto/request"
	"golang-point-of-sales-system/modules/penjualan/dto/response"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type PenjualanServiceImpl struct {
	PenjualanRepository repository.PenjualanRepository
	Validate            *validator.Validate
}

func NewPenjualanService(penjualanRepository repository.PenjualanRepository, validate *validator.Validate) PenjualanService {
	return &PenjualanServiceImpl{
		PenjualanRepository: penjualanRepository,
		Validate:            validate,
	}
}

func (service *PenjualanServiceImpl) Create(ctx context.Context, request request.PenjualanCreateRequest) response.PenjualanResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	penjualan := entity.Penjualan{
		Member_id:   request.Member_id,
		Total_item:  request.Total_item,
		Total_harga: request.Total_harga,
		Bayar:       request.Bayar,
		Diskon:      request.Diskon,
		Diterima:    request.Diterima,
	}

	penjualan = service.PenjualanRepository.Save(ctx, penjualan)

	return helper.ToPenjualanResponse(penjualan)
}

func (service *PenjualanServiceImpl) Update(ctx context.Context, request request.PenjualanUpdateRequest) response.PenjualanResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	penjualan, err := service.PenjualanRepository.FindById(ctx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	penjualan.Member_id = request.Member_id
	penjualan.Total_item = request.Total_item
	penjualan.Total_harga = request.Total_harga
	penjualan.Bayar = request.Bayar
	penjualan.Diskon = request.Diskon
	penjualan.Diterima = request.Diterima

	penjualan, err = service.PenjualanRepository.Update(ctx, penjualan)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToPenjualanResponse(penjualan)
}

func (service *PenjualanServiceImpl) Delete(ctx context.Context, penjualanId uuid.UUID) {
	penjualan, err := service.PenjualanRepository.FindById(ctx, penjualanId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.PenjualanRepository.Delete(ctx, penjualan)
}

func (service *PenjualanServiceImpl) FindById(ctx context.Context, penjualanId uuid.UUID) response.PenjualanResponse {
	penjualan, err := service.PenjualanRepository.FindById(ctx, penjualanId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToPenjualanResponse(penjualan)
}

func (service *PenjualanServiceImpl) FindAll(ctx context.Context) []response.PenjualanResponse {
	datas := service.PenjualanRepository.FindAll(ctx)
	var penjualanResponses []response.PenjualanResponse
	for _, penjualan := range datas {
		penjualanResponses = append(penjualanResponses, helper.ToPenjualanResponse(penjualan))
	}

	return penjualanResponses
}
