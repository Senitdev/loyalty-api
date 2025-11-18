package dto

type UserDTO struct {
	Email string `json:"email" binding:"required,email"`
	Role  string `json:"role" binding:"required,oneof=client merchant"`
}
