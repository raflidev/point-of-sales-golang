package service

import (
	"context"
	"golang-point-of-sales-system/modules/penjualan/dto/request"
	"golang-point-of-sales-system/modules/penjualan/dto/response"

	"github.com/google/uuid"
)

type PenjualanService interface {
	Create(ctx context.Context, request request.PenjualanCreateRequest) response.PenjualanResponse
	Update(ctx context.Context, request request.PenjualanUpdateRequest) response.PenjualanResponse
	Delete(ctx context.Context, dataId uuid.UUID)
	FindById(ctx context.Context, dataId uuid.UUID) response.PenjualanResponse
	FindAll(ctx context.Context) []response.PenjualanResponse
}
