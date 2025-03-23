package service

import (
	"context"
	"golang-point-of-sales-system/modules/products/dto/request"
	"golang-point-of-sales-system/modules/products/dto/response"

	"github.com/google/uuid"
)

type ProductService interface {
	Create(ctx context.Context, request request.ProductCreateRequest) response.ProductResponse
	Update(ctx context.Context, request request.ProductUpdateRequest) response.ProductResponse
	Delete(ctx context.Context, productId uuid.UUID)
	FindById(ctx context.Context, productId uuid.UUID) response.ProductResponse
	FindAll(ctx context.Context) []response.ProductResponse
}
