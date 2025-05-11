package request

import "github.com/google/uuid"

type PembelianCreateRequest struct {
	Supplier_id uuid.UUID `validate:"required" json:"supplier_id"`
	Total_item  int       `validate:"required" json:"total_item"`
	Total_harga int       `validate:"required" json:"total_harga"`
	Diskon      int       `validate:"required" json:"diskon"`
	Bayar       int       `validate:"required" json:"bayar"`
}

type PembelianUpdateRequest struct {
	Id          uuid.UUID `validate:"required" json:"id"`
	Supplier_id uuid.UUID `validate:"required" json:"supplier_id"`
	Total_item  int       `validate:"required" json:"total_item"`
	Total_harga int       `validate:"required" json:"total_harga"`
	Diskon      int       `validate:"required" json:"diskon"`
	Bayar       int       `validate:"required" json:"bayar"`
}
