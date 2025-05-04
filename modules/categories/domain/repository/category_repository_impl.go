package repository

import (
	"context"
	"golang-point-of-sales-system/modules/categories/domain/entity"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		DB: db,
	}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, category entity.Category) entity.Category {
	// Generate a new UUID for the category
	category.Id = uuid.New()
	result := repository.DB.Create(&category)
	if result.Error != nil {
		log.Println(result.Error)
	} else {
		log.Println("Category created successfully")
	}

	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, category entity.Category) (entity.Category, error) {
	// First check if the category exists
	var existingCategory entity.Category
	result := repository.DB.First(&existingCategory, "id = ?", category.Id)
	if result.Error != nil {
		return entity.Category{}, result.Error
	}

	// Update the category using Updates to only update non-zero fields
	result = repository.DB.Model(&existingCategory).Updates(category)
	if result.Error != nil {
		return entity.Category{}, result.Error
	}

	return existingCategory, nil
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, category entity.Category) {
	result := repository.DB.Delete(&category)
	if result.Error != nil {
		log.Println(result.Error)
	} else {
		log.Println("Category deleted successfully")
	}
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, categoryId uuid.UUID) (entity.Category, error) {
	var category entity.Category
	result := repository.DB.First(&category, "id = ?", categoryId)
	if result.Error != nil {
		return entity.Category{}, result.Error
	}
	return category, nil
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context) []entity.Category {
	var categories []entity.Category
	result := repository.DB.Find(&categories)
	if result.Error != nil {
		log.Println(result.Error)
		return nil
	}
	return categories
}
