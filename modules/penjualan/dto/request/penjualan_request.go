package request

import "github.com/google/uuid"

type PenjualanCreateRequest struct {
	Member_id   uuid.UUID `validate:"required" json:"member_id"`
	User_id     uuid.UUID `validate:"required" json:"user_id"`
	Total_item  int       `validate:"required" json:"total_item"`
	Total_harga int       `validate:"required" json:"total_harga"`
	Diskon      int       `validate:"required" json:"diskon"`
	Bayar       int       `validate:"required" json:"bayar"`
	Diterima    int       `validate:"required" json:"diterima"`
}

type PenjualanUpdateRequest struct {
	Id          uuid.UUID `validate:"required" json:"id"`
	Member_id   uuid.UUID `validate:"required" json:"member_id"`
	User_id     uuid.UUID `validate:"required" json:"user_id"`
	Total_item  int       `validate:"required" json:"total_item"`
	Total_harga int       `validate:"required" json:"total_harga"`
	Diskon      int       `validate:"required" json:"diskon"`
	Bayar       int       `validate:"required" json:"bayar"`
	Diterima    int       `validate:"required" json:"diterima"`
}
