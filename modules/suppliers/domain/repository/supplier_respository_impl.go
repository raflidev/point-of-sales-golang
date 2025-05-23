package repository

import (
	"context"
	"errors"
	"golang-point-of-sales-system/modules/suppliers/domain/entity"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SupplierRepositoryImpl struct {
	DB *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) SupplierRepository {
	return &SupplierRepositoryImpl{
		DB: db,
	}
}

func (repository *SupplierRepositoryImpl) Save(ctx context.Context, supplier entity.Supplier) entity.Supplier {
	if supplier.Id == uuid.Nil {
		supplier.Id = uuid.New()
	}
	result := repository.DB.Create(&supplier)
	if result.Error != nil {
		log.Println(result.Error)
	} else {
		log.Println("Supplier created successfully")
	}
	return supplier
}

func (repository *SupplierRepositoryImpl) Update(ctx context.Context, product entity.Supplier) (entity.Supplier, error) {
	// First check if the product exists
	var existingProduct entity.Supplier
	result := repository.DB.First(&existingProduct, "id = ?", product.Id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.Supplier{}, errors.New("product not found")
		}
		return entity.Supplier{}, result.Error
	}

	// Update the product using Updates to only update non-zero fields
	result = repository.DB.Model(&existingProduct).Updates(product)
	if result.Error != nil {
		return entity.Supplier{}, result.Error
	}

	return existingProduct, nil
}

func (repository *SupplierRepositoryImpl) Delete(ctx context.Context, product entity.Supplier) {
	result := repository.DB.Delete(&product)
	if result.Error != nil {
		log.Println(result.Error)
	} else {
		log.Println("Product deleted successfully")
	}

}

func (repository *SupplierRepositoryImpl) FindById(ctx context.Context, supplierId uuid.UUID) (entity.Supplier, error) {
	var supplier entity.Supplier
	result := repository.DB.First(&supplier, "id = ?", supplierId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.Supplier{}, result.Error
		}
		log.Println(result.Error)
		return entity.Supplier{}, result.Error
	}

	return supplier, nil
}

func (repository *SupplierRepositoryImpl) FindAll(ctx context.Context) []entity.Supplier {
	result := []entity.Supplier{}
	repository.DB.Find(&result)
	if len(result) == 0 {
		log.Println("No products found")
		return nil
	}
	return result
}
