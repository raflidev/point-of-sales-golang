package service

import (
	"context"
	"golang-point-of-sales-system/modules/users/dto/request"
	"golang-point-of-sales-system/modules/users/dto/response"

	"github.com/google/uuid"
)

type UserService interface {
	Create(ctx context.Context, request request.UserCreateRequest) response.UserResponse
	Update(ctx context.Context, request request.UserUpdateRequest) response.UserResponse
	Delete(ctx context.Context, userId uuid.UUID)
	FindById(ctx context.Context, userId uuid.UUID) response.UserResponse
	FindAll(ctx context.Context) []response.UserResponse
	Login(ctx context.Context, request request.UserLoginRequest) response.UserResponse
	ChangePassword(ctx context.Context, request request.UserChangePasswordRequest) response.UserResponse
}
