package service

import (
	"context"
	"golang-point-of-sales-system/exception"
	"golang-point-of-sales-system/helper"
	"golang-point-of-sales-system/modules/categories/domain/entity"
	"golang-point-of-sales-system/modules/categories/domain/repository"
	"golang-point-of-sales-system/modules/categories/dto/request"
	"golang-point-of-sales-system/modules/categories/dto/response"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request request.CategoryCreateRequest) response.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	category := entity.Category{
		Nama_kategori: request.Nama_kategori,
	}

	category = service.CategoryRepository.Save(ctx, category)

	return helper.ToCategoryResponse(category)

}

func (service *CategoryServiceImpl) Update(ctx context.Context, request request.CategoryUpdateRequest) response.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	category, err := service.CategoryRepository.FindById(ctx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category.Nama_kategori = request.Nama_kategori

	category, err = service.CategoryRepository.Update(ctx, category)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId uuid.UUID) {
	category, err := service.CategoryRepository.FindById(ctx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId uuid.UUID) response.CategoryResponse {
	category, err := service.CategoryRepository.FindById(ctx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []response.CategoryResponse {
	categories := service.CategoryRepository.FindAll(ctx)
	var categoryResponses []response.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, helper.ToCategoryResponse(category))
	}

	return categoryResponses
}
