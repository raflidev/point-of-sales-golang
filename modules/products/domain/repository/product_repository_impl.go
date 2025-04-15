package repository

import (
	"context"
	"errors"
	"golang-point-of-sales-system/modules/products/domain/entity"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		DB: db,
	}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, product entity.Product) entity.Product {
	result := repository.DB.Create(&product)
	if result.Error != nil {
		log.Println(result.Error)
	} else {
		log.Println("Product created successfully")
	}

	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, product entity.Product) (entity.Product, error) {
	// First check if the product exists
	var existingProduct entity.Product
	result := repository.DB.First(&existingProduct, "id = ?", product.Id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.Product{}, errors.New("product not found")
		}
		return entity.Product{}, result.Error
	}

	// Update the product using Updates to only update non-zero fields
	result = repository.DB.Model(&existingProduct).Updates(product)
	if result.Error != nil {
		return entity.Product{}, result.Error
	}

	return existingProduct, nil
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, product entity.Product) {
	result := repository.DB.Delete(&product)
	if result.Error != nil {
		log.Println(result.Error)
	} else {
		log.Println("Product deleted successfully")
	}

}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, productId uuid.UUID) (entity.Product, error) {
	var product entity.Product
	result := repository.DB.First(&product, "id = ?", productId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.Product{}, result.Error
		}
		log.Println(result.Error)
		return entity.Product{}, result.Error
	}

	return product, nil
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context) []entity.Product {
	result := []entity.Product{}
	repository.DB.Find(&result)
	if len(result) == 0 {
		log.Println("No products found")
		return nil
	}
	return result
}
