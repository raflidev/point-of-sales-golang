package entity

import (
	"time"

	"github.com/google/uuid"
)

type Supplier struct {
	Id       uuid.UUID `gorm:"primary_key;column:id;type:uuid;default:uuid_generate_v4()"`
	Nama     string    `gorm:"column:nama;type:varchar(50);unique;not null"`
	Alamat   string    `gorm:"column:alamat;type:varchar(50);not null"`
	Telepon  string    `gorm:"column:telepon;type:varchar(50);not null"`
	CreateAt time.Time `gorm:"column:create_at;autoCreateTime;default:now()"`
	UpdateAt time.Time `gorm:"column:update_at;autoCreateTime;default:now()"`
}
