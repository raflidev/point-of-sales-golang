package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `gorm:"primary_key;column:id;type:uuid;default:uuid_generate_v4()"`
	Nama     string    `gorm:"column:nama;type:varchar(50);not null"`
	Email    string    `gorm:"column:email;type:varchar(50);unique;not null"`
	Password string    `gorm:"column:password;type:varchar(50);not null"`
	Foto     string    `gorm:"column:foto;type:text;not null"`
	Role     string    `gorm:"column:role;type:varchar(50);not null"`
	CreateAt time.Time `gorm:"column:create_at;autoCreateTime;default:now()"`
	UpdateAt time.Time `gorm:"column:update_at;autoCreateTime;default:now()"`
}
