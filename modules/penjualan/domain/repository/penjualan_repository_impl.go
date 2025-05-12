package repository

import (
	"context"
	"golang-point-of-sales-system/modules/penjualan/domain/entity"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PenjualanRepositoryImpl struct {
	DB *gorm.DB
}

func NewPenjualanRepository(db *gorm.DB) PenjualanRepository {
	return &PenjualanRepositoryImpl{
		DB: db,
	}
}

func (repository *PenjualanRepositoryImpl) Save(ctx context.Context, penjualan entity.Penjualan) entity.Penjualan {
	penjualan.Id = uuid.New()

	result := repository.DB.Create(&penjualan)
	if result.Error != nil {
		log.Println(result.Error)
	} else {
		log.Println("Data created successfully")
	}

	return penjualan
}

func (repository *PenjualanRepositoryImpl) Update(ctx context.Context, penjualan entity.Penjualan) (entity.Penjualan, error) {
	var existingData entity.Penjualan
	result := repository.DB.First(&existingData, "id = ?", penjualan.Id)
	if result.Error != nil {
		return entity.Penjualan{}, result.Error
	}

	result = repository.DB.Model(&existingData).Updates(penjualan)
	if result.Error != nil {
		return entity.Penjualan{}, result.Error
	}

	return existingData, nil
}

func (repository *PenjualanRepositoryImpl) Delete(ctx context.Context, penjualan entity.Penjualan) {
	result := repository.DB.Delete(&penjualan)
	if result.Error != nil {
		log.Println(result.Error)
	} else {
		log.Println("Data deleted successfully")
	}
}

func (repository *PenjualanRepositoryImpl) FindById(ctx context.Context, penjualanId uuid.UUID) (entity.Penjualan, error) {
	var data entity.Penjualan
	result := repository.DB.First(&data, "id = ?", penjualanId)
	if result.Error != nil {
		return entity.Penjualan{}, result.Error
	}
	return data, nil
}

func (repository *PenjualanRepositoryImpl) FindAll(ctx context.Context) []entity.Penjualan {
	var datas []entity.Penjualan
	result := repository.DB.Find(&datas)
	if result.Error != nil {
		log.Println(result.Error)
		return nil
	}
	return datas
}
