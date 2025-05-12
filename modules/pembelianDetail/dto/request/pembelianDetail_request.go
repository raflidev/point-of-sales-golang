package request

import "github.com/google/uuid"

type PembelianDetailCreateRequest struct {
	Pembelian_id uuid.UUID `validate:"required" json:"pembelian_id"`
	Produk_id    uuid.UUID `validate:"required" json:"produk_id"`
	Harga_beli   int       `validate:"required" json:"harga_beli"`
	Jumlah       int       `validate:"required" json:"jumlah"`
	Subtotal     int       `validate:"required" json:"subtotal"`
}

type PembelianDetailUpdateRequest struct {
	Id           uuid.UUID `validate:"required" json:"id"`
	Pembelian_id uuid.UUID `validate:"required" json:"pembelian_id"`
	Produk_id    uuid.UUID `validate:"required" json:"produk_id"`
	Harga_beli   int       `validate:"required" json:"harga_beli"`
	Jumlah       int       `validate:"required" json:"jumlah"`
	Subtotal     int       `validate:"required" json:"subtotal"`
}
