package service

import (
	"context"
	"golang-point-of-sales-system/modules/pembelian/dto/request"
	"golang-point-of-sales-system/modules/pembelian/dto/response"

	"github.com/google/uuid"
)

type PembelianService interface {
	Create(ctx context.Context, request request.PembelianCreateRequest) response.PembelianResponse
	Update(ctx context.Context, request request.PembelianUpdateRequest) response.PembelianResponse
	Delete(ctx context.Context, pembeliId uuid.UUID)
	FindById(ctx context.Context, pembeliId uuid.UUID) response.PembelianResponse
	FindAll(ctx context.Context) []response.PembelianResponse
}
