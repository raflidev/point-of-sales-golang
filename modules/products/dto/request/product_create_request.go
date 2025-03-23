package request

import "github.com/google/uuid"

type ProductCreateRequest struct {
	Kode_produk string `validate:"required" json:"kode_produk"`
	Nama_produk string `validate:"required" json:"nama_produk"`
	Merk        string `validate:"required" json:"merk"`
	Harga_beli  int    `validate:"required" json:"harga_beli"`
	Harga_jual  int    `validate:"required" json:"harga_jual"`
	Stok        int    `validate:"required" json:"stok"`
}

type ProductUpdateRequest struct {
	Id          uuid.UUID `validate:"required" json:"id"`
	Kode_produk string    `validate:"required" json:"kode_produk"`
	Nama_produk string    `validate:"required" json:"nama_produk"`
	Merk        string    `validate:"required" json:"merk"`
	Harga_beli  int       `validate:"required" json:"harga_beli"`
	Harga_jual  int       `validate:"required" json:"harga_jual"`
	Stok        int       `validate:"required" json:"stok"`
}
