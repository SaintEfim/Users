package entity

import "github.com/google/uuid"

type UserEntity struct {
	Id   uuid.UUID `json:"id" binding:"required"`
	Name string    `json:"name" binding:"required"`
}
