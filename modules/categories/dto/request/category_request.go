package request

import "github.com/google/uuid"

type CategoryCreateRequest struct {
	Nama_kategori string `validate:"required" json:"nama_kategori"`
}

type CategoryUpdateRequest struct {
	Id            uuid.UUID `validate:"required" json:"id"`
	Nama_kategori string    `validate:"required" json:"nama_kategori"`
}
