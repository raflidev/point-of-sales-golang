package service

import (
	"context"
	"golang-point-of-sales-system/helper"
	"golang-point-of-sales-system/modules/users/domain/entity"
	"golang-point-of-sales-system/modules/users/domain/repository"
	"golang-point-of-sales-system/modules/users/dto/request"
	"golang-point-of-sales-system/modules/users/dto/response"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request request.UserCreateRequest) response.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	user := entity.User{
		Nama:     request.Nama,
		Email:    request.Email,
		Password: request.Password,
		Foto:     request.Foto,
		Role:     request.Role,
	}

	user = service.UserRepository.Save(ctx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request request.UserUpdateRequest) response.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	user, err := service.UserRepository.FindById(ctx, request.Id)
	if err != nil {
		panic(err)
	}

	user.Nama = request.Nama
	user.Email = request.Email
	user.Foto = request.Foto
	user.Role = request.Role

	user, err = service.UserRepository.Update(ctx, user)
	if err != nil {
		panic(err)
	}

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId uuid.UUID) {
	user, err := service.UserRepository.FindById(ctx, userId)
	if err != nil {
		panic(err)
	}

	service.UserRepository.Delete(ctx, user)
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId uuid.UUID) response.UserResponse {
	user, err := service.UserRepository.FindById(ctx, userId)
	if err != nil {
		panic(err)
	}

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []response.UserResponse {
	users := service.UserRepository.FindAll(ctx)
	var userResponses []response.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, helper.ToUserResponse(user))
	}
	return userResponses
}

func (service *UserServiceImpl) Login(ctx context.Context, request request.UserLoginRequest) response.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	user, err := service.UserRepository.FindByEmail(ctx, request.Email)
	if err != nil {
		panic(err)
	}

	if user.Password != request.Password {
		panic("password salah")
	}

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) ChangePassword(ctx context.Context, request request.UserChangePasswordRequest) response.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	user, err := service.UserRepository.FindByEmail(ctx, request.Email)
	if err != nil {
		panic(err)
	}

	if user.Password != request.NewPass {
		panic("password lama salah")
	}

	user.Password = request.NewPass

	user, err = service.UserRepository.Update(ctx, user)
	if err != nil {
		panic(err)
	}

	return helper.ToUserResponse(user)
}
