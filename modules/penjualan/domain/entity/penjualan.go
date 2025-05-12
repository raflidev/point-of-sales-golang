package entity

import (
	"time"

	"github.com/google/uuid"
)

type Penjualan struct {
	Id          uuid.UUID `gorm:"primary_key;column:id;type:uuid;default:uuid_generate_v4()"`
	Member_id   uuid.UUID `gorm:"column:member_id;type:uuid;not null"`
	User_id     uuid.UUID `gorm:"column:user_id;type:uuid;not null"`
	Total_item  int       `gorm:"column:total_item;type:integer;not null"`
	Total_harga int       `gorm:"column:total_harga;type:integer;not null"`
	Diskon      int       `gorm:"column:diskon;type:integer;not null"`
	Bayar       int       `gorm:"column:bayar;type:integer;not null"`
	Diterima    int       `gorm:"column:diterima;type:integer;not null"`
	CreateAt    time.Time `gorm:"column:create_at;autoCreateTime;default:now()"`
	UpdateAt    time.Time `gorm:"column:update_at;autoCreateTime;default:now()"`
}
