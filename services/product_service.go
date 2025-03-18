package services

import (
	"context"
	"golang-point-of-sales-system/model/web"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest)
	Update(ctx context.Context)
	Delete(ctx context.Context)
	FindById(ctx context.Context)
	FindAll(ctx context.Context)
}
