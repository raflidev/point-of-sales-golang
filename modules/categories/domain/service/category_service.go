package service

import (
	"context"
	"golang-point-of-sales-system/modules/categories/dto/request"
	"golang-point-of-sales-system/modules/categories/dto/response"

	"github.com/google/uuid"
)

type CategoryService interface {
	Create(ctx context.Context, request request.CategoryCreateRequest) response.CategoryResponse
	Update(ctx context.Context, request request.CategoryUpdateRequest) response.CategoryResponse
	Delete(ctx context.Context, categoryId uuid.UUID)
	FindById(ctx context.Context, categoryId uuid.UUID) response.CategoryResponse
	FindAll(ctx context.Context) []response.CategoryResponse
}
