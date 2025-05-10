package service

import (
	"context"
	"golang-point-of-sales-system/modules/members/dto/request"
	"golang-point-of-sales-system/modules/members/dto/response"

	"github.com/google/uuid"
)

type MemberService interface {
	Create(ctx context.Context, request request.MemberCreateRequest) response.MemberResponse
	Update(ctx context.Context, request request.MemberUpdateRequest) response.MemberResponse
	Delete(ctx context.Context, memberId uuid.UUID)
	FindById(ctx context.Context, memberId uuid.UUID) response.MemberResponse
	FindAll(ctx context.Context) []response.MemberResponse
}
