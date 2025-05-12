package response

import "github.com/google/uuid"

type PembelianDetailResponse struct {
	Id           uuid.UUID `json:"id"`
	Pembelian_id uuid.UUID `json:"pembelian_id"`
	Produk_id    uuid.UUID `json:"produk_id"`
	Harga_beli   int       `json:"harga_beli"`
	Jumlah       int       `json:"jumlah"`
	Subtotal     int       `json:"subtotal"`
}
