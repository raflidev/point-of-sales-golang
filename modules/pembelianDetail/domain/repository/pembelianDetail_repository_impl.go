package repository

import (
	"context"
	"golang-point-of-sales-system/modules/pembelianDetail/domain/entity"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PembelianDetailRepositoryImpl struct {
	DB *gorm.DB
}

func NewPembelianDetailRepository(db *gorm.DB) PembelianDetailRepository {
	return &PembelianDetailRepositoryImpl{
		DB: db,
	}
}

func (repository *PembelianDetailRepositoryImpl) Save(ctx context.Context, pembelianDetail entity.PembelianDetail) entity.PembelianDetail {
	pembelianDetail.Id = uuid.New()

	result := repository.DB.Create(&pembelianDetail)
	if result.Error != nil {
		log.Println(result.Error)
	} else {
		log.Println("Data created successfully")
	}

	return pembelianDetail
}

func (repository *PembelianDetailRepositoryImpl) Update(ctx context.Context, pembelianDetail entity.PembelianDetail) (entity.PembelianDetail, error) {
	var existingData entity.PembelianDetail
	result := repository.DB.First(&existingData, "id = ?", pembelianDetail.Id)
	if result.Error != nil {
		return entity.PembelianDetail{}, result.Error
	}

	result = repository.DB.Model(&existingData).Updates(pembelianDetail)
	if result.Error != nil {
		return entity.PembelianDetail{}, result.Error
	}

	return existingData, nil
}

func (repository *PembelianDetailRepositoryImpl) Delete(ctx context.Context, pembelianDetail entity.PembelianDetail) {
	result := repository.DB.Delete(&pembelianDetail)
	if result.Error != nil {
		log.Println(result.Error)
	} else {
		log.Println("Data deleted successfully")
	}
}

func (repository *PembelianDetailRepositoryImpl) FindById(ctx context.Context, pembelianDetailId uuid.UUID) (entity.PembelianDetail, error) {
	var data entity.PembelianDetail
	result := repository.DB.First(&data, "id = ?", pembelianDetailId)
	if result.Error != nil {
		return entity.PembelianDetail{}, result.Error
	}
	return data, nil
}

func (repository *PembelianDetailRepositoryImpl) FindAll(ctx context.Context) []entity.PembelianDetail {
	var datas []entity.PembelianDetail
	result := repository.DB.Find(&datas)
	if result.Error != nil {
		log.Println(result.Error)
		return nil
	}
	return datas
}
