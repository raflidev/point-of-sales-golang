package repository

import (
	"context"
	"golang-point-of-sales-system/modules/pembelianDetail/domain/entity"

	"github.com/google/uuid"
)

type PembelianDetailRepository interface {
	Save(ctx context.Context, pembelianDetail entity.PembelianDetail) entity.PembelianDetail
	Update(ctx context.Context, pembelianDetail entity.PembelianDetail) (entity.PembelianDetail, error)
	Delete(ctx context.Context, pembelianDetail entity.PembelianDetail)
	FindById(ctx context.Context, pembelianDetailId uuid.UUID) (entity.PembelianDetail, error)
	FindAll(ctx context.Context) []entity.PembelianDetail
}
