package entity

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id          uuid.UUID `gorm:"primary_key;column:id;type:uuid;default:uuid_generate_v4()"`
	Kategori_Id uuid.UUID `gorm:"column:produk_id;type:uuid;not null"`
	Kode_produk string    `gorm:"column:kode_produk;type:varchar(50);unique;not null"`
	Nama_produk string    `gorm:"column:nama_produk;type:varchar(50);not null"`
	Merk        string    `gorm:"column:merk;type:varchar(50);not null"`
	Harga_beli  int       `gorm:"column:harga_beli;type:int;not null"`
	Harga_jual  int       `gorm:"column:harga_jual;type:int;not null"`
	Stok        int       `gorm:"column:stok;type:int;not null"`
	CreateAt    time.Time `gorm:"column:create_at;autoCreateTime;default:now()"`
	UpdateAt    time.Time `gorm:"column:update_at;autoCreateTime;default:now()"`
}

func (Product) TableName() string {
	return "product"
}
