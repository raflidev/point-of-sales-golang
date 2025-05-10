package repository

import (
	"context"
	"golang-point-of-sales-system/modules/members/domain/entity"

	"github.com/google/uuid"
)

type MemberRepository interface {
	Save(ctx context.Context, member entity.Member) entity.Member
	Update(ctx context.Context, member entity.Member) (entity.Member, error)
	Delete(ctx context.Context, member entity.Member)
	FindById(ctx context.Context, memberId uuid.UUID) (entity.Member, error)
	FindAll(ctx context.Context) []entity.Member
}
