package repository

import (
	"context"
	"golang-point-of-sales-system/modules/penjualan/domain/entity"

	"github.com/google/uuid"
)

type PenjualanRepository interface {
	Save(ctx context.Context, penjualan entity.Penjualan) entity.Penjualan
	Update(ctx context.Context, penjualan entity.Penjualan) (entity.Penjualan, error)
	Delete(ctx context.Context, penjualan entity.Penjualan)
	FindById(ctx context.Context, penjualanId uuid.UUID) (entity.Penjualan, error)
	FindAll(ctx context.Context) []entity.Penjualan
}
