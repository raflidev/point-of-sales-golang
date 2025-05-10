package repository

import (
	"context"
	"golang-point-of-sales-system/modules/members/domain/entity"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MemberRepositoryImpl struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) MemberRepository {
	return &MemberRepositoryImpl{
		db: db,
	}
}

func (repository *MemberRepositoryImpl) Save(ctx context.Context, member entity.Member) entity.Member {
	// Generate a new UUID for the member
	member.Id = uuid.New()
	result := repository.db.Create(&member)
	if result.Error != nil {
		log.Println(result.Error)
	} else {
		log.Println("Member created successfully")
	}

	return member
}

func (repository *MemberRepositoryImpl) Update(ctx context.Context, member entity.Member) (entity.Member, error) {
	// First check if the member exists
	var existingMember entity.Member
	result := repository.db.First(&existingMember, "id = ?", member.Id)
	if result.Error != nil {
		return entity.Member{}, result.Error
	}

	// Update the member using Updates to only update non-zero fields
	result = repository.db.Model(&existingMember).Updates(member)
	if result.Error != nil {
		return entity.Member{}, result.Error
	}

	return existingMember, nil
}

func (repository *MemberRepositoryImpl) Delete(ctx context.Context, member entity.Member) {
	result := repository.db.Delete(&member)
	if result.Error != nil {
		log.Println(result.Error)
	} else {
		log.Println("Member deleted successfully")
	}
}

func (repository *MemberRepositoryImpl) FindById(ctx context.Context, memberId uuid.UUID) (entity.Member, error) {
	var member entity.Member
	result := repository.db.First(&member, "id = ?", memberId)
	if result.Error != nil {
		return entity.Member{}, result.Error
	}
	return member, nil
}

func (repository *MemberRepositoryImpl) FindAll(ctx context.Context) []entity.Member {
	var members []entity.Member
	result := repository.db.Find(&members)
	if result.Error != nil {
		log.Println(result.Error)
	} else {
		log.Println("Members retrieved successfully")
	}
	return members
}
