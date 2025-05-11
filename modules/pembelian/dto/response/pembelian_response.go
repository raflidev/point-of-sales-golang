package response

import "github.com/google/uuid"

type PembelianResponse struct {
	Id          uuid.UUID `json:"id"`
	Supplier_id uuid.UUID `json:"supplier_id"`
	Total_item  int       `json:"total_item"`
	Total_harga int       `json:"total_harga"`
	Diskon      int       `json:"diskon"`
	Bayar       int       `json:"bayar"`
}
