package repository

import (
	"context"
	"errors"
	"golang-point-of-sales-system/modules/users/domain/entity"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, user entity.User) entity.User {
	if user.Id == uuid.Nil {
		user.Id = uuid.New()
	}
	result := repository.db.Create(&user)
	if result.Error != nil {
		log.Println(result.Error)
	} else {
		log.Println("User created successfully")
	}
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, user entity.User) (entity.User, error) {
	// First check if the user exists
	var existingUser entity.User
	result := repository.db.First(&existingUser, "id = ?", user.Id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.User{}, errors.New("user not found")
		}
		return entity.User{}, result.Error
	}

	// Update the user using Updates to only update non-zero fields
	result = repository.db.Model(&existingUser).Updates(user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return existingUser, nil
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, user entity.User) {
	result := repository.db.Delete(&user)
	if result.Error != nil {
		log.Println(result.Error)
	} else {
		log.Println("User deleted successfully")
	}
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, userId uuid.UUID) (entity.User, error) {
	var user entity.User
	result := repository.db.First(&user, "id = ?", userId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.User{}, errors.New("user not found")
		}
		return entity.User{}, result.Error
	}
	return user, nil
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User
	result := repository.db.First(&user, "email = ?", email)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entity.User{}, errors.New("user not found")
		}
		return entity.User{}, result.Error
	}
	return user, nil
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context) []entity.User {
	var users []entity.User
	result := repository.db.Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return users
}
