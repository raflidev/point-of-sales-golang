package repository

import (
	"context"
	"golang-point-of-sales-system/modules/products/domain/entity"

	"github.com/google/uuid"
)

type ProductRepository interface {
	Save(ctx context.Context, product entity.Product) entity.Product
	Update(ctx context.Context, product entity.Product) (entity.Product, error)
	Delete(ctx context.Context, product entity.Product)
	FindById(ctx context.Context, productId uuid.UUID) (entity.Product, error)
	FindAll(ctx context.Context) []entity.Product
}
