package dto

import "github.com/google/uuid"

type UserDto struct {
	Id   uuid.UUID `json:"id" binding:"required"`
	Name string    `json:"name" binding:"required"`
}
