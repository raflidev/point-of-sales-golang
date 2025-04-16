package request

import "github.com/google/uuid"

type SupplierCreateRequest struct {
	Nama    string `validate:"required" json:"nama"`
	Alamat  string `validate:"required" json:"alamat"`
	Telepon string `validate:"required" json:"telepon"`
}

type SupplierUpdateRequest struct {
	Id      uuid.UUID `validate:"required" json:"id"`
	Nama    string    `validate:"required" json:"nama"`
	Alamat  string    `validate:"required" json:"alamat"`
	Telepon string    `validate:"required" json:"telepon"`
}
