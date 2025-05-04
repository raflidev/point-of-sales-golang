package response

import "github.com/google/uuid"

type CategoryResponse struct {
	Id            uuid.UUID `json:"id"`
	Nama_kategori string    `json:"nama_kategori"`
}
