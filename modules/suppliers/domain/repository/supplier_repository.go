package repository

import (
	"context"
	"golang-point-of-sales-system/modules/suppliers/domain/entity"

	"github.com/google/uuid"
)

type SupplierRepository interface {
	Save(ctx context.Context, supplier entity.Supplier) entity.Supplier
	Update(ctx context.Context, supplier entity.Supplier) (entity.Supplier, error)
	Delete(ctx context.Context, supplier entity.Supplier)
	FindById(ctx context.Context, SupplierId uuid.UUID) (entity.Supplier, error)
	FindAll(ctx context.Context) []entity.Supplier
}
