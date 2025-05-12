package entity

import (
	"time"

	"github.com/google/uuid"
)

type PembelianDetail struct {
	Id           uuid.UUID `gorm:"primary_key;column:id;type:uuid;default:uuid_generate_v4()"`
	Pembelian_id uuid.UUID `gorm:"column:pembelian_id;type:uuid;not null"`
	Produk_id    uuid.UUID `gorm:"column:produk_id;type:uuid;not null"`
	Harga_beli   int       `gorm:"column:harga_beli;type:integer;not null"`
	Jumlah       int       `gorm:"column:jumlah;type:integer;not null"`
	Subtotal     int       `gorm:"column:subtotal;type:integer;not null"`
	CreateAt     time.Time `gorm:"column:create_at;autoCreateTime;default:now()"`
	UpdateAt     time.Time `gorm:"column:update_at;autoCreateTime;default:now()"`
}
