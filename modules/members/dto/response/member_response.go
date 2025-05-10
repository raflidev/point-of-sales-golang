package response

import "github.com/google/uuid"

type MemberResponse struct {
	Id          uuid.UUID `json:"id"`
	Kode_member string    `json:"kode_member"`
	Nama        string    `json:"nama"`
	Telepon     string    `json:"telepon"`
	Alamat      string    `json:"alamat"`
	Keterangan  string    `json:"keterangan"`
}
