package service

import (
	"context"
	"golang-point-of-sales-system/modules/suppliers/dto/request"
	"golang-point-of-sales-system/modules/suppliers/dto/response"

	"github.com/google/uuid"
)

type SupplierService interface {
	Create(ctx context.Context, request request.SupplierCreateRequest) response.SupplierResponse
	Update(ctx context.Context, request request.SupplierUpdateRequest) response.SupplierResponse
	Delete(ctx context.Context, productId uuid.UUID)
	FindById(ctx context.Context, productId uuid.UUID) response.SupplierResponse
	FindAll(ctx context.Context) []response.SupplierResponse
}
