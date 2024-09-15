package _interface

import (
	. "UsersService/internal/models/entity"
)

type UserRepository interface {
	Get() ([]*UserEntity, error)
	GetOneByID(id string) (*UserEntity, error)
	Create(user *UserEntity) error
	Delete(id string) error
	Update(id string, user *UserEntity) error
}
