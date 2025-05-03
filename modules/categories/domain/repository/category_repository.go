package repository

import (
	"context"
	"golang-point-of-sales-system/modules/categories/domain/entity"

	"github.com/google/uuid"
)

type CategoryRepository interface {
	Save(ctx context.Context, category entity.Category) entity.Category
	Update(ctx context.Context, category entity.Category) (entity.Category, error)
	Delete(ctx context.Context, category entity.Category)
	FindById(ctx context.Context, categoryId uuid.UUID) (entity.Category, error)
	FindAll(ctx context.Context) []entity.Category
}
