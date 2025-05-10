package request

import "github.com/google/uuid"

type MemberCreateRequest struct {
	Kode_member string `validate:"required" json:"kode_member"`
	Nama        string `validate:"required" json:"nama"`
	Telepon     string `validate:"required" json:"telepon"`
	Alamat      string `validate:"required" json:"alamat"`
	Keterangan  string `json:"keterangan"`
}

type MemberUpdateRequest struct {
	Id          uuid.UUID `validate:"required" json:"id"`
	Kode_member string    `validate:"required" json:"kode_member"`
	Nama        string    `validate:"required" json:"nama"`
	Telepon     string    `validate:"required" json:"telepon"`
	Alamat      string    `validate:"required" json:"alamat"`
	Keterangan  string    `json:"keterangan"`
}
