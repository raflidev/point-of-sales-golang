package response

import "github.com/google/uuid"

type PenjualanResponse struct {
	Id          uuid.UUID `json:"id"`
	Member_id   uuid.UUID `json:"member_id"`
	User_id     uuid.UUID `json:"user_id"`
	Total_item  int       `json:"total_item"`
	Total_harga int       `json:"total_harga"`
	Diskon      int       `json:"diskon"`
	Bayar       int       `json:"bayar"`
	Diterima    int       `json:"diterima"`
}
