package entity

import (
	"time"

	"github.com/google/uuid"
)

type Member struct {
	Id          uuid.UUID `gorm:"primary_key;column:id;type:uuid;default:uuid_generate_v4()"`
	Kode_member string    `gorm:"column:kode_member;type:varchar(255);not null"`
	Nama        string    `gorm:"column:nama;type:varchar(255);not null"`
	Telepon     string    `gorm:"column:telepon;type:varchar(255);not null"`
	Alamat      string    `gorm:"column:alamat;type:text;not null"`
	Keterangan  string    `gorm:"column:keterangan;type:varchar(255);not null"`
	CreateAt    time.Time `gorm:"column:create_at;autoCreateTime;default:now()"`
	UpdateAt    time.Time `gorm:"column:update_at;autoCreateTime;default:now()"`
}
