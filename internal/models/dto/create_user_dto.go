package dto

type CreateUserDto struct {
	Name string `json:"name" binding:"required"`
}
