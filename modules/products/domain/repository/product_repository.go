package repository

import (
	"context"
	"golang-point-of-sales-system/modules/products/domain/entity"
)

type ProductRepository interface {
	Save(ctx context.Context, product entity.Product) entity.Product
	Update(ctx context.Context, product entity.Product) entity.Product
	Delete(ctx context.Context, product entity.Product)
	FindById(ctx context.Context, productId int) (entity.Product, error)
	FindAll(ctx context.Context) []entity.Product
}
