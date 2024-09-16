package dto

type UpdateUserDto struct {
	Name string `json:"name" binding:"required"`
}
