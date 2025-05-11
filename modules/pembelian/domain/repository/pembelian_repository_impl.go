package repository

import (
	"context"
	"golang-point-of-sales-system/modules/pembelian/domain/entity"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PembelianRepositoryImpl struct {
	DB *gorm.DB
}

func NewPembelianRepository(db *gorm.DB) PembelianRepository {
	return &PembelianRepositoryImpl{
		DB: db,
	}
}

func (repository *PembelianRepositoryImpl) Save(ctx context.Context, pembelian entity.Pembelian) entity.Pembelian {
	pembelian.Id = uuid.New()
	result := repository.DB.Create(&pembelian)
	if result.Error != nil {
		log.Println(result.Error)
	} else {
		log.Println("data created successfully")
	}

	return pembelian
}

func (repository *PembelianRepositoryImpl) Update(ctx context.Context, pembelian entity.Pembelian) (entity.Pembelian, error) {
	// First check if the category exists
	var existingPembelian entity.Pembelian
	result := repository.DB.First(&existingPembelian, "id = ?", pembelian.Id)
	if result.Error != nil {
		return entity.Pembelian{}, result.Error
	}

	// Update the Pembelian using Updates to only update non-zero fields
	result = repository.DB.Model(&existingPembelian).Updates(pembelian)
	if result.Error != nil {
		return entity.Pembelian{}, result.Error
	}

	return existingPembelian, nil
}

func (repository *PembelianRepositoryImpl) Delete(ctx context.Context, pembelian entity.Pembelian) {
	result := repository.DB.Delete(&pembelian)
	if result.Error != nil {
		log.Println(result.Error)
	} else {
		log.Println("data deleted successfully")
	}
}

func (repository *PembelianRepositoryImpl) FindById(ctx context.Context, pembelianId uuid.UUID) (entity.Pembelian, error) {
	var pembelian entity.Pembelian
	result := repository.DB.First(&pembelian, "id = ?", pembelianId)
	if result.Error != nil {
		return entity.Pembelian{}, result.Error
	}
	return pembelian, nil
}

func (repository *PembelianRepositoryImpl) FindAll(ctx context.Context) []entity.Pembelian {
	var pembelis []entity.Pembelian
	result := repository.DB.Find(&pembelis)
	if result.Error != nil {
		log.Println(result.Error)
		return nil
	}
	return pembelis
}
