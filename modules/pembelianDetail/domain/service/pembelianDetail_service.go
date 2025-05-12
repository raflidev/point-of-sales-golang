package service

import (
	"context"
	"golang-point-of-sales-system/modules/pembelianDetail/dto/request"
	"golang-point-of-sales-system/modules/pembelianDetail/dto/response"

	"github.com/google/uuid"
)

type PembelianDetailService interface {
	Create(ctx context.Context, request request.PembelianDetailCreateRequest) response.PembelianDetailResponse
	Update(ctx context.Context, request request.PembelianDetailUpdateRequest) response.PembelianDetailResponse
	Delete(ctx context.Context, dataId uuid.UUID)
	FindById(ctx context.Context, dataId uuid.UUID) response.PembelianDetailResponse
	FindAll(ctx context.Context) []response.PembelianDetailResponse
}
