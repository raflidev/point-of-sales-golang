package response

import "github.com/google/uuid"

type UserResponse struct {
	Id    uuid.UUID `json:"uuid"`
	Nama  string    `json:"nama"`
	Email string    `json:"email"`
	Foto  string    `json:"foto"`
	Role  string    `json:"role"`
}
