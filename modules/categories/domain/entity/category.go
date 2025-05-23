package entity

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	Id            uuid.UUID `gorm:"primary_key;column:id;type:uuid;default:uuid_generate_v4()"`
	Nama_kategori string    `gorm:"column:nama_kategori;type:varchar(50);not null"`
	CreateAt      time.Time `gorm:"column:create_at;autoCreateTime;default:now()"`
	UpdateAt      time.Time `gorm:"column:update_at;autoCreateTime;default:now()"`
}

func (Category) TableName() string {
	return "kategori"
}
