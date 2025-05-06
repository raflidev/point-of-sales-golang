package repository

import (
	"context"
	"golang-point-of-sales-system/modules/users/domain/entity"

	"github.com/google/uuid"
)

type UserRepository interface {
	Save(ctx context.Context, user entity.User) entity.User
	Update(ctx context.Context, user entity.User) (entity.User, error)
	Delete(ctx context.Context, user entity.User)
	FindById(ctx context.Context, userId uuid.UUID) (entity.User, error)
	FindByEmail(ctx context.Context, email string) (entity.User, error)
	FindAll(ctx context.Context) []entity.User
}
