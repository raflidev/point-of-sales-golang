package response

import "github.com/google/uuid"

type SupplierResponse struct {
	Id      uuid.UUID `json:"id"`
	Nama    string    `json:"nama"`
	Alamat  string    `json:"alamat"`
	Telepon string    `json:"telepon"`
}
