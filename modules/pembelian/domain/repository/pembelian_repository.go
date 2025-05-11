package repository

import (
	"context"
	"golang-point-of-sales-system/modules/pembelian/domain/entity"

	"github.com/google/uuid"
)

type PembelianRepository interface {
	Save(ctx context.Context, pembelian entity.Pembelian) entity.Pembelian
	Update(ctx context.Context, pembelian entity.Pembelian) (entity.Pembelian, error)
	Delete(ctx context.Context, pembelian entity.Pembelian)
	FindById(ctx context.Context, pembelianId uuid.UUID) (entity.Pembelian, error)
	FindAll(ctx context.Context) []entity.Pembelian
}
