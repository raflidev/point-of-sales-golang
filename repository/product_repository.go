package repository

import (
	"context"
	"golang-point-of-sales-system/model/domain"

	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sqlx.Tx, product domain.Product) domain.Product
	Update(ctx context.Context, tx *sqlx.Tx, product domain.Product) domain.Product
	Delete(ctx context.Context, tx *sqlx.Tx, product domain.Product)
	FindById(ctx context.Context, tx *sqlx.Tx, productId int) (domain.Product, error)
	FindAll(ctx context.Context, tx *sqlx.Tx) []domain.Product
}
