package request

import "github.com/google/uuid"

type UserCreateRequest struct {
	Nama     string `validate:"required" json:"nama"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
	Foto     string `validate:"required" json:"foto"`
	Role     string `validate:"required" json:"role"`
}

type UserUpdateRequest struct {
	Id       uuid.UUID `validate:"required" json:"id"`
	Nama     string    `validate:"required" json:"nama"`
	Email    string    `validate:"required,email" json:"email"`
	Password string    `validate:"required" json:"password"`
	Foto     string    `validate:"required" json:"foto"`
	Role     string    `validate:"required" json:"role"`
}

type UserLoginRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type UserChangePasswordRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
	NewPass  string `validate:"required" json:"new_pass"`
}
